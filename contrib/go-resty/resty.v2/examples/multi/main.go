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

	// call google

	googleopt := new(resty.Options)

	err = config.UnmarshalWithPath("app.client.resty.google", googleopt)
	if err != nil {
		logger.Errorf(err.Error())
	}

	healthIntegrator := health.NewDefaultIntegrator()

	cligoogle := resty.NewClientWithOptions(ctx, googleopt, healthIntegrator.Register)
	reqgoogle := cligoogle.R()

	var respgoogle *r.Response

	respgoogle, err = reqgoogle.Get("/")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if respgoogle != nil {
		logger.Infof(respgoogle.String())
	}

	// call acom

	acomopt := new(resty.Options)

	err = config.UnmarshalWithPath("app.client.resty.acom", acomopt)
	if err != nil {
		logger.Errorf(err.Error())
	}

	cliacom := resty.NewClientWithOptions(ctx, acomopt)
	reqacom := cliacom.R()

	var respacom *r.Response

	respacom, err = reqacom.Get("/")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if respacom != nil {
		logger.Infof(respacom.String())
	}
}
