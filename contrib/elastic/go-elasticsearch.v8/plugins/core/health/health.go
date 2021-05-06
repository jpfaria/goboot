package health

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/elastic/go-elasticsearch/v8"
)

type Health struct {
	options *Options
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}
func (i *Health) Register(ctx context.Context, client *elasticsearch.Client) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating elasticsearch in health")

	checker := NewChecker(client)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("elasticsearch successfully integrated in health")

	return nil
}
