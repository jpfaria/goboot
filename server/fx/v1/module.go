package fx

import (
	"context"
	"sync"

	"github.com/b2wdigital/goignite/v2/cobra/v1"
	contextfx "github.com/b2wdigital/goignite/v2/context/fx/v1"
	"github.com/b2wdigital/goignite/v2/server"
	c "github.com/spf13/cobra"
	"go.uber.org/fx"
)

const (
	ServersGroupKey = "_gi_server_servers_"
)

type srvParams struct {
	fx.In
	Servers []server.Server `group:"_gi_server_servers_"`
}

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {

		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				func(ctx context.Context, p srvParams) error {

					return cobra.Run(
						&c.Command{
							Run: func(cmd *c.Command, args []string) {
								server.Serve(ctx, p.Servers...)
							},
						},
					)

				},
			),
		)
	})

	return options
}
