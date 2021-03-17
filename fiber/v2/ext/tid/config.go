package tid

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	root    = fiber.ExtRoot + ".tid"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable tid middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
