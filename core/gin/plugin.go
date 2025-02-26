//go:build gin && !admin

package core

import (
	"fmt"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
)

var Plugin = new(_plugin)

type _plugin struct{}

// Register 注册插件
func (c *_plugin) Register(group *gin.RouterGroup, plugin interfaces.GinPlugin) {
	fmt.Printf("[%s]插件注册开始!\n", plugin.RouterPath())
	group = group.Group(plugin.RouterPath())
	plugin.Register(group)
	fmt.Printf("[%s]插件注册成功!\n", plugin.RouterPath())
}

// Initialization 初始化插件
func (c *_plugin) Initialization(public, private *gin.RouterGroup) {}
