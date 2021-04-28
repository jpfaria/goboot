package compressor

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func Register(ctx context.Context) []grpc.ServerOption {

	logger := log.FromContext(ctx)
	logger.Debug("compressor successfully enabled in grpc server")

	err := gzip.SetLevel(Level())
	if err != nil {
		logger.Fatalf("could not set level: %s", err.Error())
	}

	return nil
}
