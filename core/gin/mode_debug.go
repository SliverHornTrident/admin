//go:build gin && debug

package core

import "github.com/gin-gonic/gin"

func init() {
	gin.SetMode(gin.DebugMode)
}
