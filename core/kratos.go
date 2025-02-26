//go:build kratos

package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Kratos = new(_kratos)

type _kratos struct{}

func (c *_kratos) Name() string {
	return "[shadow][core][kratos]"
}

func (c *_kratos) Viper(viper *viper.Viper) error {

	return nil
}

func (c *_kratos) IsPanic() bool {
	return true
}

func (c *_kratos) ConfigName() string {
	return strings.Join([]string{"kratos", "yaml"}, ".")
}

func (c *_kratos) Initialization(ctx context.Context) error {
	{
		options := []grpc.ServerOption{
			grpc.Middleware(recovery.Recovery()),
		}
		if address := global.KratosConfig.Http.Address(); address != "" {
			options = append(options, grpc.Address(address))
		}
		if global.KratosConfig.Grpc.Timeout != 0 {
			options = append(options, grpc.Timeout(global.KratosConfig.Grpc.Timeout))
		}
		global.KratosGrpc = grpc.NewServer(options...)
	} // grpc
	{
		opts := []http.ServerOption{
			http.Middleware(
				recovery.Recovery(),
			),
		}
		if address := global.KratosConfig.Http.Address(); address != "" {
			opts = append(opts, http.Address(address))
		}
		if global.KratosConfig.Http.Timeout != 0 {
			opts = append(opts, http.Timeout(global.KratosConfig.Http.Timeout))
		}
		global.KratosHttp = http.NewServer(opts...)
	} // http

	app := kratos.New(
		kratos.ID(global.KratosConfig.Id),
		kratos.Name(global.KratosConfig.Name),
		kratos.Version(global.KratosConfig.Version),
		kratos.Server(global.KratosGrpc, global.KratosHttp),
	)
	err := app.Run()
	if err != nil {
		return errors.Wrap(err, "app run failed!")
	}
	return nil
}
