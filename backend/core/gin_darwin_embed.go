//go:build gin && darwin && embed

package core

import (
	"context"
	internal "github.com/SliverHornTrident/shadow/core/gin"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var _ interfaces.Corer = (*_gin)(nil)

var Gin = new(_gin)

type _gin struct{}

func (s *_gin) Name() string {
	return "[shadow][core][gin][engine][darwin][embed]"
}

func (s *_gin) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Gin", &global.GinConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (s *_gin) IsPanic() bool {
	return true
}

func (s *_gin) ConfigName() string {
	return strings.Join([]string{"gin", gin.Mode(), "yaml"}, ".")
}

func (s *_gin) Initialization(ctx context.Context) error {
	global.Gin = gin.Default()
	server := endless.NewServer(global.GinConfig.Address(), global.Gin)
	server.ReadHeaderTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.MaxHeaderBytes = 1 << 20
	internal.Static.Set()
	internal.Router.Initialization(ctx)
	zap.L().Info("server run success on ", zap.String("address", global.GinConfig.Address()))
	zap.L().Error(server.ListenAndServe().Error())
	return nil
}
