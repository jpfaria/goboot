package status

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/b2wdigital/goignite/v2/core/rest/response"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, app *f.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	statusRoute := getRoute()

	logger.Tracef("configuring status router on %s in fiber", statusRoute)

	app.Get(statusRoute, func(c *f.Ctx) error {
		return fiber.JSON(c, http.StatusOK, response.NewResourceStatus(), nil)
	})

	logger.Debugf("status router configured on %s in fiber", statusRoute)

	return nil
}
