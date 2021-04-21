package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1"
	newrelic2 "github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1/plugins/contrib/newrelic/go-agent.v3"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
)

func main() {

	config.Load()
	logrus.NewLogger()

	mongo.NewConn(context.Background(), newrelic2.Register)
}
