package newrelic

import (
	"context"

	newrelic "github.com/b2wdigital/goignite/v2/contrib/newrelic/go-agent.v3"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.DialOption {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("newrelic interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
		grpc.WithChainStreamInterceptor(nrgrpc.StreamClientInterceptor),
	}

}
