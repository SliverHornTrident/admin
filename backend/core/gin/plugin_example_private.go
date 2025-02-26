//go:build gin

package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var PluginExamplePrivate = new(pluginExamplePrivate)

type pluginExamplePrivate struct{}

func (e *pluginExamplePrivate) Register(group *gin.RouterGroup) {
	fmt.Println("测试插件=>私有路由组")
}

func (e *pluginExamplePrivate) RouterPath() string {
	fmt.Println("测试插件=>私有路由组")
	return ""
}
