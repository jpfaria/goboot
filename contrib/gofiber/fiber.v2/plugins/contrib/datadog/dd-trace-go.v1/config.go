package datadog

import (
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = fiber.PluginsRoot + ".datadog.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable fiber integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
