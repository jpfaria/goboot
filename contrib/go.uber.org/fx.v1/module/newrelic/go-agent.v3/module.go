package newrelic

import (
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	newrelic "github.com/b2wdigital/goignite/v2/contrib/newrelic/go-agent.v3"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				newrelic.NewApplication,
			),
		)
	})

	return options
}
