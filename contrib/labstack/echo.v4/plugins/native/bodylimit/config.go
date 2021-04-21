package bodylimit

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = echo.ExtRoot + ".bodylimit.enabled"
	size    = echo.ExtRoot + ".bodylimit.size"
)

func init() {
	config.Add(enabled, true, "enable/disable body limit middleware")
	config.Add(size, "8M", "body limit size")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetSize() string {
	return config.String(size)
}
