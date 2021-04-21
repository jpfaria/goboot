package echo_pprof_v1

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = echo.ExtRoot + ".pprof.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable pprof integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
