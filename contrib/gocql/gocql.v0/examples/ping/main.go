package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/gocql/gocql.v0"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
)

func main() {

	config.Load()

	logrus.NewLogger()

	session, err := gocql.NewSession(context.Background())
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.Query("void").Exec()
	if err != nil {
		panic(err)
	}

}
