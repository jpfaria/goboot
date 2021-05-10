package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1"
	newrelic "github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1/plugins/contrib/newrelic/go-agent.v3"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
)

func main() {

	config.Load()

	logrus.NewLogger()

	conn, err := mongo.NewConn(context.Background(), newrelic.Register)
	if err != nil {
		log.Panic(err)
	}

	err = conn.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

}
