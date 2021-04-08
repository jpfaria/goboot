package datadog

import (
	"sync"

	datadog "github.com/b2wdigital/goignite/v2/contrib/datadog/dd-trace-go.v1"
	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var optOnce sync.Once

func OptionsModule() fx.Option {
	options := fx.Options()

	optOnce.Do(func() {
		options = fx.Options(
			fx.Provide(
				datadog.DefaultOptions,
			),
		)
	})

	return options
}

var tracerOnce sync.Once

func TracerModule() fx.Option {
	options := fx.Options()

	tracerOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			OptionsModule(),
			fx.Invoke(
				datadog.NewTracer,
			),
		)
	})

	return options
}

var profileOnce sync.Once

func ProfileModule() fx.Option {
	options := fx.Options()

	profileOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			OptionsModule(),
			fx.Invoke(
				datadog.NewProfiler,
			),
		)
	})

	return options
}
