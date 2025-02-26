//go:build gin

package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var PluginExamplePublic = new(pluginExamplePublic)

type pluginExamplePublic struct{}

func (e *pluginExamplePublic) Register(group *gin.RouterGroup) {
	fmt.Println("测试插件=>公开路由组")
}

func (e *pluginExamplePublic) RouterPath() string {
	fmt.Println("测试插件=>公开路由组")
	return ""
}
