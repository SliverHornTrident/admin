package interfaces

import (
	"context"
	"github.com/spf13/viper"
)

type Corer interface {
	// Name 核心名称
	Name() string
	// Viper viper
	Viper(viper *viper.Viper) error
	// IsPanic 是否panic退出程序
	IsPanic() bool
	// ConfigName 核心对应配置名
	ConfigName() string
	// Initialization 初始化
	Initialization(ctx context.Context) error
}
