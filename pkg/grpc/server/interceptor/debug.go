package interceptor

import (
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func DebugStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		wrapper := &recvWrapper{stream}
		err := handler(srv, wrapper)
		log.Debugf("invoke server method=%s duration=%s error=%v", info.FullMethod,
			time.Since(start), err)
		return err
	}
}

func DebugUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		r, err := handler(ctx, req)
		log.Debugf("invoke server method=%s duration=%s error=%v response=%v", info.FullMethod,
			time.Since(start), err, r)
		return r, err
	}
}

type recvWrapper struct {
	grpc.ServerStream
}
