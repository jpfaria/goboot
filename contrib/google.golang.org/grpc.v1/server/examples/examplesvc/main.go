package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/server"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
)

func main() {

	ctx := context.Background()

	config.Load()

	// start logrus
	logrus.NewLogger()

	srv := server.NewServer(ctx)

	RegisterExampleServer(srv.Server(), &Service{})

	srv.Serve(ctx)
}

type Service struct {
}

func (h *Service) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {

	logger := log.FromContext(ctx)

	logger.Infof(request.Message)

	return &TestResponse{Message: "hello world"}, nil
}
