package datadog

import (
	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = mongo.PluginsRoot + ".datadog"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
