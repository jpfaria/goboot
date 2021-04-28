package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/go-resty/resty.v2"
	"github.com/b2wdigital/goignite/v2/contrib/go-resty/resty.v2/plugins/core/health"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
	r "github.com/go-resty/resty/v2"
)

func main() {

	var err error

	config.Load()

	ctx := context.Background()

	logrus.NewLogger()

	logger := log.FromContext(ctx)

	options := health.Options{
		Name:        "Google Inc",
		Host:        "http://google.com",
		Endpoint:    "/status",
		Enabled:     true,
		Description: "Search Engine",
		Required:    true,
	}

	healthIntegrator := health.NewIntegrator(&options)

	client := resty.NewClientWithOptions(ctx, &resty.Options{}, healthIntegrator.Register)
	request := client.R().EnableTrace()

	var resp *r.Response
	resp, err = request.Get("http://google.com")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if resp != nil {
		logger.Infof(resp.String())
	}
}
