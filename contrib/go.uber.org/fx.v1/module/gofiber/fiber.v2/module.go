package fiber

import (
	"context"
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	serverfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/core/server"
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/server"
	f "github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type params struct {
	fx.In
	Exts []fiber.Ext `optional:"true"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				func(ctx context.Context, p params) *fiber.Server {
					return fiber.NewServer(ctx, p.Exts...)
				},
				func(srv *fiber.Server) *f.App {
					return srv.App()
				},
			),
			fx.Provide(
				fx.Annotated{
					Group: serverfx.ServersGroupKey,
					Target: func(srv *fiber.Server) server.Server {
						return srv
					},
				},
			),
		)
	})

	return options
}
