package newrelic

import (
	ginats "github.com/b2wdigital/goignite/v2/contrib/nats-io/nats.go.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = ginats.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
