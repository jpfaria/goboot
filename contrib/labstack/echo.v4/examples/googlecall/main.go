package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/go-resty/resty.v2"
	"github.com/b2wdigital/goignite/v2/contrib/go.uber.org/zap.v1"
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
	r "github.com/go-resty/resty/v2"
	e "github.com/labstack/echo/v4"
)

const Endpoint = "app.endpoint.google"

func init() {
	config.Add(Endpoint, "/google", "google endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Google string
		}
	}
}

type Response struct {
	Message string
}

type Handler struct {
	client *r.Client
}

func NewHandler(client *r.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) Get(c e.Context) (err error) {

	log := log.FromContext(c.Request().Context())

	request := h.client.R().EnableTrace()

	_, err = request.Get("http://google.com")
	if err != nil {
		log.Fatalf(err.Error())
	}

	resp := Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		log.Errorf(err.Error())
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

	zap.NewLogger()

	info.AppName = "google"

	srv := echo.NewServer(ctx,
		cors2.Register,
		requestid2.Register,
		gzip2.Register,
		logger.Register,
		status.Register,
		health.Register)

	// instance.AddErrorAdvice(customErrors.InvalidPayload, 400)

	o := resty.OptionsBuilder.
		Host("http://www.google.com").
		Build()

	client := resty.NewClientWithOptions(ctx, &o)

	handler := NewHandler(client)
	srv.Instance().GET(c.App.Endpoint.Google, handler.Get)

	srv.Serve(ctx)
}
