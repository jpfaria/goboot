package compressor

import (
	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/client"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root  = client.PluginsRoot + ".compressor"
	level = root + ".level"
)

func init() {
	config.Add(level, -1, "sets gzip level")
}

func Level() int {
	return config.Int(level)
}
