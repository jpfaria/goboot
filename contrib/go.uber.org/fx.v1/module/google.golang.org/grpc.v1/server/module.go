package server

import (
	"context"
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	serverfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/core/server"
	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/server"
	s "github.com/b2wdigital/goignite/v2/core/server"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type params struct {
	fx.In
	Plugins []server.Plugin `optional:"true"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *server.Server {
					return server.NewServer(ctx, p.Plugins...)
				},
				func(srv *server.Server) *grpc.Server {
					return srv.Server()
				},
				func(srv *server.Server) grpc.ServiceRegistrar {
					return srv.ServiceRegistrar()
				},
				fx.Annotated{
					Group: serverfx.ServersGroupKey,
					Target: func(srv *server.Server) s.Server {
						return srv
					},
				},
			),
		)

	})

	return options
}
