package newrelic

import (
	"context"

	newrelic "github.com/b2wdigital/goignite/v2/contrib/newrelic/go-agent.v3"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in newrelic")

	client.AddHook(nrredis.NewHook(client.Options()))

	logger.Debug("redis successfully integrated in newrelic")

	return nil
}
