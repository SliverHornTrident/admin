//go:build zap

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/core/internal"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var _ interfaces.Corer = (*_zap)(nil)

var Zap = new(_zap)

type _zap struct{}

func (c *_zap) Name() string {
	return "[shadow][core][zap]"
}

func (c *_zap) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Zap", &global.ZapConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_zap) IsPanic() bool {
	return false
}

func (c *_zap) ConfigName() string {
	return strings.Join([]string{"zap", gin.Mode(), "yaml"}, ".")
}

func (c *_zap) Initialization(ctx context.Context) error {
	if global.ZapConfig.InOutputFile {
		_ = os.Mkdir(global.ZapConfig.Director, os.ModePerm)
	}
	levels := global.ZapConfig.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...)) // 初始化 zap.Logger
	if global.ZapConfig.ShowLine {              // 判断是否显示行
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger) // logger 注册到全局, 通过 zap.L() 调用日志组件
	return nil
}
