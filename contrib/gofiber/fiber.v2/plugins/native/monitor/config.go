package monitor

import (
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = fiber.ExtRoot + ".monitor.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable monitor middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
