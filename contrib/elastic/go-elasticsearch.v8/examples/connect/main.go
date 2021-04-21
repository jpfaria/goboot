package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/elastic/go-elasticsearch.v8"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
)

func main() {

	config.Load()

	logrus.NewLogger()

	client, err := elasticsearch.NewClient(context.Background())
	if err != nil {
		log.Panic(err)
	}

	ping, err := client.Ping(client.Ping.WithPretty())
	if err != nil {
		log.Panic(err)
	}

	log.Infof("status: %v", ping.StatusCode)

}
