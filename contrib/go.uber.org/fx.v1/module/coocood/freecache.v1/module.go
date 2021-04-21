package freecache

import (
	"sync"

	"github.com/b2wdigital/goignite/v2/contrib/coocood/freecache.v1"
	contextfx "github.com/b2wdigital/goignite/v2/contrib/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			fx.Invoke(
				freecache.NewCache,
			),
		)
	})

	return options
}
