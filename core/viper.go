//go:build viper

package core

import (
	"context"
	"flag"
	"fmt"
	"github.com/SliverHornTrident/shadow/constant"
	"github.com/SliverHornTrident/shadow/core/internal"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Viper.pwd, _ = os.Getwd()
	flag.StringVar(&Viper.configs, "configs", "", "choose configs folder.")
	flag.Parse()
	if Viper.configs == "" { // 判断命令行参数是否为空
		env := os.Getenv(constant.ViperEnv)
		if env == "" { // 判断环境变量是否为空
			Viper.configs = constant.ViperConfigs
			fmt.Printf("您正在使用默认值,config的路径为%s\n", filepath.Join(Viper.pwd, Viper.configs))
		} else {
			Viper.configs = env
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", constant.ViperEnv, Viper.configs)
		}
	} else {
		fmt.Printf("您正在使用命令行的 -configs 参数传递的值,configs的路径为%s\n", Viper.configs)
	}
}

var _ interfaces.Corer = (*_viper)(nil)

var Viper = new(_viper)

type _viper struct {
	pwd     string
	configs string
	cores   map[string]interfaces.Corer
}

func (c *_viper) Name() string {
	return "[shadow][core][viper]"
}

// Viper .
// 优先级: 命令行 > 环境变量 > 默认值
func (c *_viper) Viper(viper *viper.Viper) (err error) {
	return nil
}

func (c *_viper) IsPanic() bool {
	return true
}

func (c *_viper) ConfigName() string {
	return ""
}

func (c *_viper) Initialization(ctx context.Context) error {
	v := viper.New()
	v.AddConfigPath(c.configs)
	entries, err := internal.Viper.Files(c.configs)
	if err != nil {
		return err
	}
	for i := 0; i < len(entries); i++ {
		var configName, configMode, configType string
		filename := entries[i].Name()
		names := strings.Split(filename, ".")
		length := len(names)
		if length < 3 {
			continue
		}
		configName = strings.Join(names[:length-2], ".")
		configMode = names[length-2]
		configType = names[length-1]
		if configMode != gin.Mode() {
			continue
		}
		configName = strings.Join([]string{configName, configMode}, ".")
		v.SetConfigName(configName)
		v.SetConfigType(configType)
		err = v.MergeInConfig()
		if err != nil {
			return errors.Wrapf(err, "[filename:%s]读取配置文件失败!", filename)
		}
		v.WatchConfig()
		v.OnConfigChange(func(event fsnotify.Event) {
			name := filepath.Base(event.Name)
			core, ok := c.cores[name]
			if !ok {
				zap.L().Info(fmt.Sprintf("[viper][filename:%s]配置文件更新!", event.Name), zap.Error(err), zap.String("business", "viper"))
				return
			}
			var bytes []byte
			bytes, err = os.ReadFile(event.Name)
			if err != nil {
				zap.L().Error(fmt.Sprintf("[viper][filename:%s]配置文件读取失败!", event.Name), zap.Error(err), zap.String("business", "viper"))
				return
			}
			err = global.Viper.ReadConfig(strings.NewReader(string(bytes)))
			if err != nil {
				zap.L().Error(fmt.Sprintf("[viper][filename:%s]配置文件读取失败!", event.Name), zap.Error(err), zap.String("business", "viper"))
				return
			}
			err = core.Viper(global.Viper)
			if err != nil {
				zap.L().Error(fmt.Sprintf("[viper][filename:%s]配置文件更新失败!", event.Name), zap.Error(err), zap.String("business", "viper"))
			}
			err = core.Initialization(ctx)
			if err != nil {
				zap.L().Error(fmt.Sprintf("[viper][filename:%s]配置文件更新初始化失败!", event.Name), zap.Error(err), zap.String("business", "viper"))
				return
			}
			zap.L().Info(fmt.Sprintf("[viper][filename:%s]配置文件更新重启组件成功!", event.Name), zap.Error(err), zap.String("business", "viper"))
		})
	}
	global.Viper = v
	return nil
}
