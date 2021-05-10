package fx

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root     = "gi.fx"
	logLevel = root + ".log.level"
)

func init() {
	config.Add(logLevel, "DEBUG", "define log level")
}

func LogLevel() string {
	return config.String(logLevel)
}
