package gzip

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = echo.PluginsRoot + ".gzip.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable gzip middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
