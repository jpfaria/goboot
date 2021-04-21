package log

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("logger interceptor successfully enabled in grpc server")

	return []grpc.ServerOption{
		grpc.ChainStreamInterceptor(streamInterceptor()),
		grpc.ChainUnaryInterceptor(unaryInterceptor()),
	}
}
