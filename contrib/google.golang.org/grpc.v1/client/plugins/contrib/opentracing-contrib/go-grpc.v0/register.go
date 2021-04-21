package opentracing

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.DialOption {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("opentracing interceptor successfully enabled in grpc client")

	tracer := opentracing.GlobalTracer()

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithChainStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	}

}
