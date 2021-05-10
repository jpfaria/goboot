package logrus

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root         = "gi.log.logrus"
	redisRoot    = root + ".redis"
	redisEnabled = redisRoot + ".enabled"
	redisHost    = redisRoot + ".host"
	redisKey     = redisRoot + ".key"
	redisFormat  = redisRoot + ".format"
	redisApp     = redisRoot + ".app"
	redisPort    = redisRoot + ".port"
	redisDb      = redisRoot + ".db"
	formatter    = root + ".formatter"
)

func init() {

	config.Add(redisEnabled, false, "enable/disable redis logging")
	config.Add(redisHost, "localhost", "redis host")
	config.Add(redisKey, "", "redis key")
	config.Add(redisFormat, "v1", "redis format")
	config.Add(redisApp, "logger", "redis app")
	config.Add(redisPort, 6379, "redis port")
	config.Add(redisDb, 0, "redis db")
	config.Add(formatter, "TEXT", "formatter TEXT/JSON/AWS_CLOUD_WATCH")
}
