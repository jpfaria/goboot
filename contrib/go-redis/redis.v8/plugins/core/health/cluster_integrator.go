package health

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-redis/redis/v8"
)

type ClusterIntegrator struct {
	options *Options
}

func NewClusterIntegrate(options *Options) *ClusterIntegrator {
	return &ClusterIntegrator{options: options}
}

func NewDefaultClusterIntegrator() *ClusterIntegrator {
	o, err := DefaultOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClusterIntegrate(o)
}

func (i *ClusterIntegrator) Register(ctx context.Context, client *redis.ClusterClient) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClusterClientChecker(client)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}
