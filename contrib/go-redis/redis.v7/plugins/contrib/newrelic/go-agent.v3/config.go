package newrelic

import (
	"github.com/b2wdigital/goignite/v2/contrib/go-redis/redis.v7"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = redis.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable redis integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
