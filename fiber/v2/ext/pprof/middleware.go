package pprof

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func Middleware(ctx context.Context, app *fiber.App) error {
	if isEnabled() {
		app.Use(pprof.New())
	}

	return nil
}