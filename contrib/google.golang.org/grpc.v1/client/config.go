package client

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root                  = "gi.grpc.client"
	skuHost               = ".host"
	skuTLS                = ".tls"
	skuGzip               = ".gzip"
	skuCertFile           = ".certFile"
	skuKeyFile            = ".keyFile"
	skuCAFile             = ".caFile"
	skuHostOverwrite      = ".hostOverwrite"
	skuPort               = ".port"
	skuInsecureSkipVerify = ".insecureSkipVerify"
	PluginsRoot           = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+skuHost, "localhost", "defines sku host")
	config.Add(path+skuPort, 9091, "defines sku port")
	config.Add(path+skuTLS, true, "enable/disable sku tls")
	config.Add(path+skuGzip, true, "enable/disable sku gzip")
	config.Add(path+skuCertFile, "./cert/server.crt", "defines sku cert file")
	config.Add(path+skuKeyFile, "./cert/server.key", "defines sku key file")
	config.Add(path+skuCAFile, "./cert/server.crt", "defines sku ca file")
	config.Add(path+skuHostOverwrite, "", "defines offer host overwrite")
	config.Add(path+skuInsecureSkipVerify, true, "enable/disable sku insecure skip verify ")
}
