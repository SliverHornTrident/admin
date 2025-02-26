//go:build windows && embed

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/mattn/go-colorable"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.DefaultWriter = colorable.NewColorableStdout()
}

var Gin = new(_gin)

type _gin struct{}

func (s *_gin) Name() string {
	return "[shadow][core][gin][engine][windows][embed]"
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
	server := http.Server{
		Addr:           global.GinConfig.Address(),
		Handler:        global.Gin,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	internal.Static.Set()
	internal.Router.Initialization(ctx)
	zap.L().Info("server run success on ", zap.String("address", global.GinConfig.Address()))
	zap.L().Error(server.ListenAndServe().Error())
	return nil
}
