package tid

import (
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = fiber.PluginsRoot + ".tid"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable tid middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
