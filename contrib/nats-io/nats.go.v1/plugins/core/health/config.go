package health

import (
	ginats "github.com/b2wdigital/goignite/v2/contrib/nats-io/nats.go.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root        = ginats.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	config.Add(name, "nats", "health name")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
