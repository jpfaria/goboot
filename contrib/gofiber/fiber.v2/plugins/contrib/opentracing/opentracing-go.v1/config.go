package opentracing

import (
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = fiber.PluginsRoot + ".opentracing.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable opentracing middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
