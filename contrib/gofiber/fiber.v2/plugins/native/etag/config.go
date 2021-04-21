package etag

import (
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = fiber.ExtRoot + ".etag.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable etag middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
