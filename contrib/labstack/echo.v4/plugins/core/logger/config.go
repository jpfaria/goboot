package logger

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = echo.PluginsRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	config.Add(enabled, true, "enable/disable logging request middleware")
	config.Add(level, "INFO", "sets log level INFO/DEBUG/TRACE")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func Level() string {
	return config.String(level)
}
