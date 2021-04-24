package swagger

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = echo.PluginsRoot + ".swagger"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable swagger integration")
	config.Add(route, "/swagger/*", "define swagger metrics url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetRoute() string {
	return config.String(route)
}
