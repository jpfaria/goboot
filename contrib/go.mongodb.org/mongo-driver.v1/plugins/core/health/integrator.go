package health

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1"
	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
	m "go.mongodb.org/mongo-driver/mongo"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) *Integrator {
	return &Integrator{options: options}
}

func NewDefaultIntegrator() *Integrator {
	o, err := DefaultOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewIntegrator(o)
}

func (i *Integrator) Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {

	logger := log.WithTypeOf(*i)

	return nil, func(ctx context.Context, client *m.Client) error {

		logger.Trace("integrating mongo in health")

		checker := NewChecker(client)
		hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
		health.Add(hc)

		logger.Debug("mongo successfully integrated in health")

		return nil
	}
}
