package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"strings"
)

var Fiber = new(_fiber)

type _fiber struct{}

func (c *_fiber) Name() string {
	return "[shadow][core][gin][engine][fiber]"
}

func (c *_fiber) Viper(viper *viper.Viper) error {
	return nil
}

func (c *_fiber) IsPanic() bool {
	return true
}

func (c *_fiber) ConfigName() string {
	return strings.Join([]string{"fiber", gin.Mode(), "yaml"}, ".")
}

func (c *_fiber) Initialization(ctx context.Context) error {
	fiber.New()
	return nil
}
