package logger

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gigrpc "github.com/b2wdigital/goignite/grpc/v1/client"
)

const (
	root    = gigrpc.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
