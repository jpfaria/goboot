package main

import (
	"os"

	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
)

type AppConfig struct {
	Application struct {
		Name       string
		MyNameTest string
	}
}

func init() {
	config.Add("app.application.name", "app_test", "name of application")
	config.Add("app.application.myName", "my_name_test", "name of application")
}

func main() {

	os.Setenv("APP_APPLICATION_NAME", "app_test_env")
	os.Setenv("APP_APPLICATION_MY.NAME.TEST", "my_name_test_env")
	os.Setenv("CONF", "config.yaml")

	config.Load()
	logrus.NewLogger()

	c := AppConfig{}

	config.UnmarshalWithPath("app", &c)

	log.Info(c.Application.Name)
	log.Info(c.Application.MyNameTest)
}
