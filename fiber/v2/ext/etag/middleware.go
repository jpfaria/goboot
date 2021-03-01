package etag

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func Middleware(ctx context.Context, app *fiber.App) error {
	if isEnabled() {
		app.Use(etag.New())
	}

	return nil
}
