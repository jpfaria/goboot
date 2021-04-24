package newrelic

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root                       = echo.PluginsRoot + ".newrelic"
	enabled                    = root + ".enabled"
	middlewareRoot             = root + ".middleware"
	middlewareRequestIDEnabled = middlewareRoot + ".requestId.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
	config.Add(middlewareRequestIDEnabled, true, "enable/disable request id middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func IsEnabledRequestID() bool {
	return config.Bool(middlewareRequestIDEnabled)
}
