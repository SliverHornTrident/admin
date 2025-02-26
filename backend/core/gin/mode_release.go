//go:build release

package core

import "github.com/gin-gonic/gin"

func init() {
	gin.SetMode(gin.ReleaseMode)
}
