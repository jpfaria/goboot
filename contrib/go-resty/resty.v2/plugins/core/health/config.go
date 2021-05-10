package health

import (
	"github.com/b2wdigital/goignite/v2/contrib/go-resty/resty.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root        = resty.PluginsRoot + ".health"
	name        = root + ".name"
	host        = root + ".host"
	endpoint    = root + ".endpoint"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {

	config.Add(name, "rest api", "health name")
	config.Add(host, "", "health host")
	config.Add(endpoint, "/resource-status", "health host")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
