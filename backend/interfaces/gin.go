//go:build gin

package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Shutdown(context.Context) error
	ListenAndServe() error
}

// GinPlugin 插件模式接口化
type GinPlugin interface {
	// Register 注册路由
	Register(group *gin.RouterGroup)
	// RouterPath 用户返回注册路由
	RouterPath() string
}
