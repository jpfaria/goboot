package prometheus

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(ctx context.Context) (*chi.Config, error) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling prometheus middleware in chi")

	prometheusRoute := getRoute()

	logger.Tracef("configuring prometheus router on %s in chi", prometheusRoute)

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			promMiddleware,
		},
		Handlers: []chi.ConfigHandler{
			{
				Handler: promhttp.Handler(),
				Pattern: prometheusRoute,
			},
		},
	}, nil

}
