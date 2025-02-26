//go:build viper && embed

package internal

import (
	"github.com/SliverHornTrident/shadow/global"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func (v *viper) Embed(pwd, configs string) error {
	info, err := os.Stat(filepath.Join(pwd, configs))
	if os.IsNotExist(err) && info.IsDir() { // 文件夹不存在
		err = os.MkdirAll(filepath.Join(pwd, configs), os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "创建文件夹失败!")
		}
	}
	entries, err := global.Configs.ReadDir(configs)
	if err != nil {
		return errors.Wrap(err, "读取文件夹失败!")
	}
	for i := 0; i < len(entries); i++ {
		filename := filepath.Base(entries[i].Name())
		names := strings.Split(filename, ".")
		if len(names) < 3 {
			continue
		}
		length := len(names)
		configMode := names[length-2]
		if configMode != gin.Mode() { // 判断部署的环境和配置文件的环境是否一致
			continue
		}
		_, err = os.Stat(filepath.Join(pwd, configs, filename))
		if !os.IsNotExist(err) { // 文件存在
			continue
		}
		var body []byte
		body, err = global.Configs.ReadFile(filepath.Join(configs, filename))
		if err != nil {
			return errors.Wrap(err, "读取文件失败!")
		}
		err = os.WriteFile(filepath.Join(pwd, configs, filename), body, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "写入文件失败!")
		}
	}
	return nil
}
