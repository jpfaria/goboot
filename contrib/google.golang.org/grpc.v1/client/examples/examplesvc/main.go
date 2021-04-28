package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/client"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
)

func main() {

	ctx := context.Background()

	config.Load()

	logger := logrus.NewLogger()

	request := &TestRequest{
		Message: "mensagem da requisição",
	}

	options := client.Options{
		Host: "localhost",
		Port: 9090,
	}

	conn := client.NewClientConnWithOptions(ctx, &options)
	defer conn.Close()

	c := NewExampleClient(conn)

	test, err := c.Test(ctx, request)
	if err != nil {
		logger.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	logger.Infof(test.Message)

	logger.Infof(conn.GetState().String())
}
