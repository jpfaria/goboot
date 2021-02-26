package semaphore

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/labstack/echo/v4"
	"golang.org/x/sync/semaphore"
)

var (
	sem *semaphore.Weighted
)

func Middleware(ctx context.Context, instance *echo.Echo) error {
	if isEnabled() {
		instance.Use(Semaphore(int64(getLimit())))
	}
	return nil
}

func Semaphore(limit int64) echo.MiddlewareFunc {

	sem = semaphore.NewWeighted(limit)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			logger := gilog.FromContext(c.Request().Context())

			if !sem.TryAcquire(1) {
				logger.Errorf("the http server has reached the limit of %v simultaneous connections", limit)
				return echo.ErrServiceUnavailable
			}
			defer sem.Release(1)

			return next(c)
		}
	}
}