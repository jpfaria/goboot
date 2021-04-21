package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/plugins/core/health"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/plugins/core/logger"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/plugins/core/status"
	cors2 "github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/plugins/native/cors"
	gzip2 "github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/plugins/native/gzip"
	requestid2 "github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/plugins/native/requestid"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/info"
	"github.com/b2wdigital/goignite/v2/core/log"
	e "github.com/labstack/echo/v4"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	config.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Helloworld string
		}
	}
}

type Response struct {
	Message string
}

func Get(c e.Context) (err error) {

	l := log.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Errorf(err.Error())
	}

	return echo.JSON(c, http.StatusOK, resp, err)
}

func main() {

	config.Load()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	//logrus.NewLogger()
	//zap.NewLogger()
	//zap.log.NewLogger()

	info.AppName = "helloworld"

	srv := echo.NewServer(ctx,
		cors2.Register,
		requestid2.Register,
		gzip2.Register,
		logger.Register,
		status.Register,
		health.Register)

	srv.Instance().GET(c.App.Endpoint.Helloworld, Get)

	srv.Serve(ctx)
}
