package newrelic

import (
	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root               = chi.PluginsRoot + ".newrelic"
	enabled            = root + ".enabled"
	webResponseEnabled = root + ".webresponse.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic middleware")
	config.Add(webResponseEnabled, true, "enable/disable newrelic web response")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func isWebResponseEnabled() bool {
	return config.Bool(enabled)
}
