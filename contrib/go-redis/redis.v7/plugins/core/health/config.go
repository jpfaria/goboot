package health

import (
	"github.com/b2wdigital/goignite/v2/contrib/go-redis/redis.v7"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root        = redis.PluginsRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {

	config.Add(name, "redis", "health name")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
