package cors

import (
	"net/http"

	gifiber "github.com/b2wdigital/goignite/v2/contrib/gofiber/fiber.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/gofiber/fiber/v2"
)

const (
	root             = gifiber.PluginsRoot + ".cors"
	enabled          = root + ".enabled"
	allowOrigins     = root + ".allow.origins"
	allowHeaders     = root + ".allow.headers"
	allowMethods     = root + ".allow.methods"
	allowCredentials = root + ".allow.credentials"
	exposeHeaders    = root + ".expose.headers"
	maxAge           = root + ".maxAge"
)

func init() {
	config.Add(enabled, true, "enable/disable cors middleware")
	config.Add(allowOrigins, []string{"*"}, "cors allow origins")
	config.Add(allowHeaders, []string{fiber.HeaderOrigin, fiber.HeaderContentType, fiber.HeaderAccept},
		"cors allow headers")
	config.Add(allowMethods,
		[]string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		"cors allow methods")
	config.Add(allowCredentials, true, "cors allow credentials")
	config.Add(exposeHeaders, []string{}, "cors expose headers")
	config.Add(maxAge, 5200, "cors max age (seconds)")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getAllowOrigins() []string {
	return config.Strings(allowOrigins)
}

func getAllowMethods() []string {
	return config.Strings(allowMethods)
}

func getAllowHeaders() []string {
	return config.Strings(allowHeaders)
}

func getAllowCredentials() bool {
	return config.Bool(allowCredentials)
}

func getExposeHeaders() []string {
	return config.Strings(exposeHeaders)
}

func getMaxAge() int {
	return config.Int(maxAge)
}
