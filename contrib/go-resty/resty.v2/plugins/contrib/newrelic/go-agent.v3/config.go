package newrelic

import (
	girest "github.com/b2wdigital/goignite/v2/contrib/go-resty/resty.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	ConfigRoot = girest.ExtRoot + ".newrelic"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
