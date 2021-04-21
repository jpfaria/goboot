package freecache

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root      = "gi.freecache"
	cacheSize = root + ".cacheSize"
)

func init() {
	config.Add(cacheSize, 100*1024*1024, "The cache size will be set to 512KB at minimum")
}
