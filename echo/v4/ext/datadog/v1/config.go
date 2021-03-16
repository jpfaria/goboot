package giechodatadog

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	enabled = giecho.ExtRoot + ".datadog.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable datadog middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
