//go:build gin && test

package core

import "github.com/gin-gonic/gin"

func init() {
	gin.SetMode(gin.TestMode)
}
