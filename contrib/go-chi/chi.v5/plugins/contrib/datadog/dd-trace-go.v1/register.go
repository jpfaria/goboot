package datadog

import (
	"context"
	"net/http"

	datadog "github.com/b2wdigital/goignite/v2/contrib/datadog/dd-trace-go.v1"
	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/core/log"
	c "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi.v5"
)

func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() || !datadog.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			c.Middleware(c.WithServiceName(datadog.Service())),
		},
	}, nil
}
