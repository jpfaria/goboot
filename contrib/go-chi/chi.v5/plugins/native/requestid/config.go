package requestid

import (
	"github.com/b2wdigital/goignite/v2/contrib/go-chi/chi.v5"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	enabled = chi.ExtRoot + ".requestid.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable requestid middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
