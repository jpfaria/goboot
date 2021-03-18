package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/echo/v4"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/cors"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/gzip"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/health"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/logger"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/requestid"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/status"
	"github.com/b2wdigital/goignite/v2/info"
	"github.com/b2wdigital/goignite/v2/log"
	girest "github.com/b2wdigital/goignite/v2/resty/v2"
	"github.com/b2wdigital/goignite/v2/zap/v1"
	"github.com/go-resty/resty/v2"
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
	client *resty.Client
}

func NewHandler(client *resty.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) Get(c e.Context) (err error) {

	logger := log.FromContext(c.Request().Context())

	request := h.client.R().EnableTrace()

	_, err = request.Get("http://google.com")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	resp := Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		logger.Errorf(err.Error())
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

	srv := echo.NewDefault(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		logger.Register,
		status.Register,
		health.Register)

	// instance.AddErrorAdvice(customErrors.InvalidPayload, 400)

	o := girest.OptionsBuilder.
		Host("http://www.google.com").
		Build()

	client := girest.NewClient(ctx, &o)

	handler := NewHandler(client)
	srv.Echo().GET(c.App.Endpoint.Google, handler.Get)

	srv.Serve(ctx)
}