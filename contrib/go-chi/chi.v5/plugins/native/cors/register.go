package cors

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-chi/cors"
)

func Register(ctx context.Context) (*chi.Config, error) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling cors middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			cors.Handler(cors.Options{
				AllowedOrigins:   getAllowedOrigins(),
				AllowedMethods:   getAllowedMethods(),
				AllowedHeaders:   getAllowedHeaders(),
				AllowCredentials: getAllowedCredentials(),
				ExposedHeaders:   getExposedHeaders(),
				MaxAge:           getMaxAge(),
			}),
		},
	}, nil
}
