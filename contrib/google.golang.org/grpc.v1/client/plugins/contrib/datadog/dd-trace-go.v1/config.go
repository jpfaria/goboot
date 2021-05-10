package datadog

import (
	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/client"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = client.PluginsRoot + ".datadog"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
