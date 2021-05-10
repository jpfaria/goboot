package echo

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root              = "gi.echo"
	hideBanner        = root + ".hidebanner"
	port              = root + ".port"
	jsonPrettyEnabled = root + ".json.pretty.enabled"
	PluginsRoot       = root + ".plugins"
)

func init() {
	config.Add(hideBanner, true, "echo hide/show banner")
	config.Add(port, 8080, "Server http port")
	config.Add(jsonPrettyEnabled, false, "enable/disable json pretty response")
}

func GetJSONPrettyEnabled() bool {
	return config.Bool(jsonPrettyEnabled)
}
