package main

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/goignite/v2/contrib/go-redis/redis.v8"
	"github.com/b2wdigital/goignite/v2/contrib/go-redis/redis.v8/plugins/core/health"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	h "github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
)

func main() {

	config.Load()

	logrus.NewLogger()

	var err error

	healthIntegrator := health.NewDefaultClientIntegrator()

	_, err = redis.NewClient(context.Background(), healthIntegrator.Register)
	if err != nil {
		log.Error(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
