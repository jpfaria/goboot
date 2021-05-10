package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1"
	newrelic "github.com/b2wdigital/goignite/v2/contrib/newrelic/go-agent.v3"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil, nil
	}

	return func(ctx context.Context, options *options.ClientOptions) error {
		logger := log.FromContext(ctx)

		logger.Trace("integrating mongo in newrelic")

		options.SetMonitor(nrmongo.NewCommandMonitor(options.Monitor))

		logger.Debug("mongo successfully integrated in newrelic")

		return nil
	}, nil
}
