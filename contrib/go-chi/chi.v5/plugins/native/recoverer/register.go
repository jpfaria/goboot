package recoverer

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-chi/chi/v5/middleware"
)

func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling recoverer middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.Recoverer,
		},
	}, nil
}
