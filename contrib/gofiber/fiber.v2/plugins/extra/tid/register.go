package tid

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/info"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling tid middleware in fiber")

	app.Use(tidMiddleware())

	logger.Debug("tid middleware successfully enabled in fiber")

	return nil
}

func tidMiddleware() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) error {

		tid := c.Get("X-TID", info.AppName+"-"+uuid.NewV4().String())
		c.Context().SetUserValue("x-tid", tid)
		c.Append("X-TID", tid)

		// Continue stack
		return c.Next()
	}
}
