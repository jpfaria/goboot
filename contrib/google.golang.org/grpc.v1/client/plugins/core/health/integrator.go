package health

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
	"google.golang.org/grpc"
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

func (i *Integrator) Register(ctx context.Context, conn *grpc.ClientConn) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating grpc with health")

	checker := NewChecker(conn)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("grpc integrated on health with success")

	return nil
}
