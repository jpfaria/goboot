package giechotid

import (
	"context"

	giinfo "github.com/b2wdigital/goignite/v2/info"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("enabling tid middleware in echo")

	instance.Use(tidMiddleware())

	logger.Debug("recover tid successfully enabled in echo")

	return nil
}

// tidMiddleware returns a middleware that tid HTTP requests.
func tidMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ctx := c.Request().Context()

			tid, ok := ctx.Value("x-tid").(string)
			if !ok {
				tid = giinfo.AppName + "-" + uuid.NewV4().String()
			}

			c.Response().Header().Add("X-TID", tid)
			c.Request().WithContext(context.WithValue(ctx, "x-tid", tid))

			return next(c)
		}
	}
}
