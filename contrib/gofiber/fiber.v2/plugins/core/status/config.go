package status

import (
	"github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = fiber.PluginsRoot + ".status"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable status route")
	config.Add(route, "/resource-status", "define status url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getRoute() string {
	return config.String(route)
}
