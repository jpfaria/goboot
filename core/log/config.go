package log

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root           = "gi.log"
	ConsoleEnabled = root + ".console.enabled"
	ConsoleLevel   = root + ".console.level"
	fileRoot       = root + ".file"
	FileEnabled    = fileRoot + ".enabled"
	FileLevel      = fileRoot + ".level"
	FilePath       = fileRoot + ".path"
	FileName       = fileRoot + ".name"
	FileMaxSize    = fileRoot + ".maxsize"
	FileCompress   = fileRoot + ".compress"
	FileMaxAge     = fileRoot + ".maxage"
)

func init() {

	config.Add(ConsoleEnabled, true, "enable/disable console logging")
	config.Add(ConsoleLevel, "INFO", "console log level")
	config.Add(FileEnabled, false, "enable/disable file logging")
	config.Add(FileLevel, "INFO", "console log level")
	config.Add(FilePath, "/tmp", "log path")
	config.Add(FileName, "application.l", "log filename")
	config.Add(FileMaxSize, 100, "log file max size (MB)")
	config.Add(FileCompress, true, "log file compress")
	config.Add(FileMaxAge, 28, "log file max age (days)")

}
