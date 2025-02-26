//go:build kratos

package global

import (
	"github.com/SliverHornTrident/shadow/config"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	KratosGrpc   *grpc.Server
	KratosHttp   *http.Server
	KratosConfig config.Kratos
)
