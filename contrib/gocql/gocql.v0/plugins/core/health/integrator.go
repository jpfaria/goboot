package health

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/gocql/gocql"
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

func (i *Integrator) Register(ctx context.Context, session *gocql.Session) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating gocql in health")

	checker := NewChecker(session)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("gocql successfully integrated in health")

	return nil
}
