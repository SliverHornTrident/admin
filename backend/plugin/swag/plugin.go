package swag

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.RouterGroup) {
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (p *plugin) RouterPath() string {
	return "swagger"
}
