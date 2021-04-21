package newrelic

import (
	"github.com/b2wdigital/goignite/v2/contrib/panjf2000/ants.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = ants.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
