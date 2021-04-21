package fiber

import (
	"context"
	"strconv"

	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/gofiber/fiber/v2"
)

type Ext func(context.Context, *fiber.App) error

type Server struct {
	app     *fiber.App
	options *Options
}

func NewServer(ctx context.Context, exts ...Ext) *Server {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, options, exts...)
}

func NewServerWithOptions(ctx context.Context, options *Options, exts ...Ext) *Server {

	app := fiber.New(*options.Config)

	for _, ext := range exts {
		if err := ext(ctx, app); err != nil {
			panic(err)
		}
	}

	return &Server{app: app, options: options}
}

func (s *Server) App() *fiber.App {
	return s.app
}

func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)
	logger.Infof("starting fiber Server. https://gofiber.io/")

	addr := ":" + strconv.Itoa(s.options.Port)

	logger.Fatal(s.app.Listen(addr))
}
