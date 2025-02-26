package interfaces

import (
	"io"
	"os"
)

type Viper interface {
	// File 获取文件信息
	File(name string) (io.Reader, error)
	// Files 获取配置文件夹信息
	Files(path string) ([]os.DirEntry, error)
	// Update 更新配置文件
	Update(configs, name string, bytes []byte) error
}
