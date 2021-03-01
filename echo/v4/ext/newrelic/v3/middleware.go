package newrelic

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Middleware(ctx context.Context, instance *echo.Echo) error {

	if !isEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating echo with newrelic")

	instance.Use(nrecho.Middleware(ginewrelic.Application()))

	if isEnabledRequestID() {
		logger.Debug("enabling newrelic requestID middleware")
		instance.Use(requestIDMiddleware())
	}

	logger.Debug("echo integrated with newrelic with success")

	return nil
}

func requestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := c.Request().Context()
			txn := newrelic.FromContext(ctx)
			reqId := c.Request().Header.Get(echo.HeaderXRequestID)
			if reqId == "" {
				reqId = c.Response().Header().Get(echo.HeaderXRequestID)
			}

			txn.AddAttribute("request.id", reqId)
			return next(c)
		}
	}
}