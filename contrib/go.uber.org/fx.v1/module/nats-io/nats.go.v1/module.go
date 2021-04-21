package nats

import (
	"sync"

	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	"github.com/b2wdigital/goignite/v2/contrib/nats-io/nats.go.v1"
	"go.uber.org/fx"
)

var subsOnce sync.Once

func SubscriberModule() fx.Option {
	options := fx.Options()

	subsOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				nats.NewSubscriber,
			),
		)
	})

	return options
}

var pubOnce sync.Once

func PublisherModule() fx.Option {
	options := fx.Options()

	pubOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Provide(
				nats.NewPublisher,
			),
		)
	})

	return options
}
