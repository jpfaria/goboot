package zap

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	ConsoleFormatter = "gi.log.zap.console.formatter"
	FileFormatter    = "gi.log.zap.file.formatter"
)

func init() {

	config.Add(ConsoleFormatter, "TEXT", "formatter TEXT/JSON")
	config.Add(FileFormatter, "TEXT", "formatter TEXT/JSON")
}
