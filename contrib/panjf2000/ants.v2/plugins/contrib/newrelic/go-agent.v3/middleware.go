package newrelic

import (
	"context"

	newrelic "github.com/b2wdigital/goignite/v2/contrib/newrelic/go-agent.v3"
	"github.com/b2wdigital/goignite/v2/contrib/panjf2000/ants.v2"
	"github.com/b2wdigital/goignite/v2/core/log"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

type middleware struct {
}

func (i *middleware) Before(ctx context.Context) context.Context {

	if IsEnabled() || !newrelic.IsEnabled() {
		return ctx
	}

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("creating go routine for newrelic")

	txn := newrelic.FromContext(ctx).NewGoroutine()

	logger.Debug("goroutine for newrelic successfully created in context")

	return nr.NewContext(ctx, txn)
}

func (i *middleware) After(ctx context.Context) {

	if IsEnabled() || !newrelic.IsEnabled() {
		return
	}

}

func NewMiddleware() ants.Middleware {
	log.Trace("creating newrelic middleware for ants")
	return &middleware{}
}
