package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	cors2 "github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2/plugins/native/cors"
	etag2 "github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2/plugins/native/etag"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/info"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/goignite/v2/core/server"
	f "github.com/gofiber/fiber/v2"
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

func Get(c *f.Ctx) (err error) {

	logger := log.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func main() {

	config.Load()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	logrus.NewLogger()

	info.AppName = "helloworld"

	fiberSrv := fiber.NewServer(ctx,
		cors2.Register,
		etag2.Register)

	fiberSrv.App().Get(c.App.Endpoint.Helloworld, Get)

	fiberSrv.Serve(ctx)

	chiSrv := chi.NewServer(ctx)

	server.Serve(ctx, fiberSrv, chiSrv)
}
