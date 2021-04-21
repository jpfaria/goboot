package health

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/goignite/v2/core/rest/response"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *e.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	healthRoute := GetRoute()

	logger.Tracef("configuring health router on %s in echo", healthRoute)

	healthHandler := NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)

	logger.Debugf("health router configured on %s in echo", healthRoute)

	return nil
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c e.Context) error {

	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	resp, httpCode := response.NewHealth(ctx)

	return echo.JSON(c, httpCode, resp, nil)
}
