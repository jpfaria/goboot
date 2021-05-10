package logger

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root = "gi.log"
	impl = root + ".impl"
)

func init() {
	config.Add(impl, "LOGRUS", "defines log implementation LOGRUS/ZAP/ZEROLOG")
}

func Impl() string {
	return config.String(impl)
}
