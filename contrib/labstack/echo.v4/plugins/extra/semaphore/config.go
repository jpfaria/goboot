package semaphore

import (
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	semaphoreRoot = echo.ExtRoot + ".semaphore"
	enabled       = semaphoreRoot + ".enabled"
	limit         = semaphoreRoot + ".limit"
)

func init() {
	config.Add(enabled, true, "enable/disable semaphore middleware")
	config.Add(limit, 10000, "defines numbers for concurrent connections")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetLimit() int {
	return config.Int(limit)
}
