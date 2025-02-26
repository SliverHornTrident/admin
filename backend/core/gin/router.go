//go:build gin

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/global"
)

var Router = new(router)

type router struct{}

func (r *router) Initialization(ctx context.Context) {
	public := global.Gin.Group("")
	private := global.Gin.Group("")
	Plugin.Initialization(public, private)
	RouterPing.Initialization(public)
}
