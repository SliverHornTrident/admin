//go:build etcd

package core

import (
	"context"
	"fmt"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var _ interfaces.Corer = (*_etcd)(nil)

var Etcd = new(_etcd)

type _etcd struct{}

func (c *_etcd) Name() string {
	return "[shadow][core][gin][engine][etcd]"
}

func (c *_etcd) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Etcd", &global.EtcdConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_etcd) IsPanic() bool {
	return true
}

func (c *_etcd) ConfigName() string {
	return strings.Join([]string{"etcd", gin.Mode(), "yaml"}, ".")
}

func (c *_etcd) Initialization(ctx context.Context) error {
	client, err := etcd.New(etcd.Config{
		Endpoints:   global.EtcdConfig.Endpoints,
		DialTimeout: global.EtcdConfig.DialTimeout,
		Username:    global.EtcdConfig.Username,
		Password:    global.EtcdConfig.Password,
	})
	if err != nil {
		return errors.Wrap(err, "链接失败!")
	}
	global.Etcd = client
	return nil
}

func (c *_etcd) viper() error {
	mod := gin.Mode()
	keys := []string{
		fmt.Sprintf("/%s/configs/gorm.%s.yaml", global.GinConfig.Name, mod),
		fmt.Sprintf("/%s/configs/gorms.%s.yaml", global.GinConfig.Name, mod),
		fmt.Sprintf("/%s/configs/mongo.%s.yaml", global.GinConfig.Name, mod),
		fmt.Sprintf("/%s/configs/redis.%s.yaml", global.GinConfig.Name, mod),
		fmt.Sprintf("/%s/configs/zap.%s.yaml", global.GinConfig.Name, mod),
		fmt.Sprintf("/%s/%s_%s", global.GinConfig.Name, global.GinConfig.Name, mod),
	}
	pwd, _ := os.Getwd()
	for i := 0; i < len(keys); i++ {
		response, err := etcd.NewKV(global.Etcd).Get(context.Background(), keys[i])
		if err != nil {
			return errors.Wrapf(err, "[etcd][viper][key:%s]读取配置中心失败!", keys[i])
		}
		if len(response.Kvs) != 0 {
			filename := string(response.Kvs[0].Key) // 文件名
			content := response.Kvs[0].Value        // 文件内容
			filename, err = filepath.Rel(global.GinConfig.Etcd(), filename)
			if err != nil {
				return errors.Wrapf(err, "[etcd][viper][filename:%s]修改绝对路径失败!", filename)
			}
			err = c.persistence(filepath.Join(pwd, filename), content)
			if err != nil {
				return err
			}
		}
		go func(key string, version int64) {
			for {
				err = c.watch(pwd, key, version)
				if err != nil {
					zap.L().Error(fmt.Sprintf("%+v", err), zap.String("business", "etcd"))
				}
			}
		}(keys[i], response.Header.Revision+1)
	}
	return nil
}

func (c *_etcd) watch(pwd, key string, version int64) error {
	watcher := etcd.NewWatcher(global.Etcd)
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})
	watchResponseChan := watcher.Watch(ctx, key, etcd.WithRev(version))
	for watchResponse := range watchResponseChan {
		for j := 0; j < len(watchResponse.Events); j++ {
			switch watchResponse.Events[j].Type {
			case mvccpb.PUT:
				filename, err := filepath.Rel(global.GinConfig.Etcd(), string(watchResponse.Events[j].Kv.Key))
				if err != nil {
					return errors.Wrapf(err, "[etcd][watch][filename:%s]修改绝对路径失败!", string(watchResponse.Events[j].Kv.Key))
				}
				if filename == global.GinConfig.EtcdBinary() {
					url := string(watchResponse.Events[j].Kv.Value)
					var request *http.Request
					request, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
					if err != nil {
						return errors.Wrapf(err, "[etcd][watch][filename:%s]创建请求失败!", filename)
					}
					var response *http.Response
					response, err = http.DefaultClient.Do(request)
					if err != nil {
						return errors.Wrapf(err, "[etcd][watch][filename:%s]请求失败!", filename)
					}
					var file *os.File
					file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
					if errors.Is(err, os.ErrNotExist) {
						file, err = os.Create(filename)
						if err != nil {
							return errors.Wrapf(err, "[etcd][filename:%s]创建二进制文件失败!", filename)
						}
					}
					_, err = io.Copy(file, response.Body)
					if err != nil {
						return errors.Wrapf(err, "[etcd][watch][filename:%s]写入二进制文件失败!", filename)
					}
					_ = file.Close()
					_ = response.Body.Close()
					command := global.GinConfig.EtcdCommand()
					err = exec.Command(command).Run()
					if err != nil {
						return errors.Wrapf(err, "[etcd][watch][filename:%s][node:]重启失败!", filename)
					}
				}
				err = c.persistence(filepath.Join(pwd, filename), watchResponse.Events[j].Kv.Value)
				if err != nil {
					return errors.Wrapf(err, "[etcd][watch][filename:%s]持久化失败!", string(watchResponse.Events[j].Kv.Key))
				}
			case mvccpb.DELETE:
				return errors.New(fmt.Sprintf("[etcd][watch][key:%s]配置中心删除了配置文件!", string(watchResponse.Events[j].Kv.Key)))
			}
		}
	} // 处理kv变化事件
	return nil
}

// persistence 持久化
func (c *_etcd) persistence(filename string, content []byte) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(filename)
		if err != nil {
			return errors.Wrapf(err, "[etcd][filename:%s]创建配置文件失败!", filename)
		}
	}
	_, err = file.Write(content)
	if err != nil {
		return errors.Wrapf(err, "[etcd][filename:%s]写入配置文件失败!", filename)
	}
	return nil
}
