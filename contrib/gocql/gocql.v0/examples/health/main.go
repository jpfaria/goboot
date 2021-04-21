package main

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/goignite/v2/contrib/gocql/gocql.v0"
	h "github.com/b2wdigital/goignite/v2/contrib/gocql/gocql.v0/plugins/core/health"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
)

func main() {

	config.Load()

	logrus.NewLogger()

	integrator := h.NewDefaultIntegrator()

	session, err := gocql.NewSession(context.Background(), integrator.Register)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	all := health.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
