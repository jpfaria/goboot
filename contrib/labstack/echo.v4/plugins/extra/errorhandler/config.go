package errorhandler

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = echo.PluginsRoot + ".errorhandler.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable custom error handler")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
