//go:build gin

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"github.com/SliverHornTrident/shadow/interfaces"
	"github.com/gin-gonic/gin"
)

var (
	Gin       *gin.Engine
	Server    interfaces.Server
	GinConfig config.Gin
)
