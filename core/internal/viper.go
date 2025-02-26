//go:build viper

package internal

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"os"
)

var Viper = new(viper)

type viper struct{}

func (v *viper) Files(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "[viper][path:%s]获取配置文件夹信息失败!", path)
	}
	return entries, nil
}
func (v *viper) File(name string) (io.Reader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, errors.Wrap(err, "文件不存在!")
	}
	defer func() {
		_ = file.Close()
	}()
	all, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "读取文件流失败!")
	}
	return bytes.NewReader(all), nil
}
