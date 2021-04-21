package prometheus

import (
	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/server"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = server.ExtRoot + ".prometheus"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable prometheus")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
