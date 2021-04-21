package client

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root                          = "gi.fasthttp.client"
	name                          = root + ".name"
	noDefaultUserAgentHeader      = root + ".noDefaultUserAgentHeader"
	maxConnsPerHost               = root + ".maxConnsPerHost"
	maxConnWaitTimeout            = root + ".maxConnWaitTimeout"
	readBufferSize                = root + ".readBufferSize"
	writeBufferSize               = root + ".writeBufferSize"
	readTimeout                   = root + ".readTimeout"
	writeTimeout                  = root + ".writeTimeout"
	maxIdleConnDuration           = root + ".maxIdleConnDuration"
	disableHeaderNamesNormalizing = root + ".disableHeaderNamesNormalizing"
	dialDualStack                 = root + ".dialDualStack"
	maxResponseBodySize           = root + ".maxResponseBodySize"
	maxIdemponentCallAttempts     = root + ".maxIdemponentCallAttempts"
)

func init() {

	config.Add(name, "", "used in User-Agent request header")
	config.Add(noDefaultUserAgentHeader, false, "User-Agent header to be excluded from the Request")
	config.Add(maxConnsPerHost, 512, "the maximum number of concurrent connections")
	config.Add(readBufferSize, 0, "per-connection buffer size for responses' reading")
	config.Add(writeBufferSize, 0, "per-connection buffer size for requests' writing")
	config.Add(maxConnWaitTimeout, 0*time.Second, "maximum amount of time to wait for a connection to be free")
	config.Add(readTimeout, 0*time.Second, "maximum duration for full response reading (including body)")
	config.Add(writeTimeout, 0*time.Second, "maximum duration for full request writing (including body)")
	config.Add(maxIdleConnDuration, 10*time.Second, "the default duration before idle keep-alive")
	config.Add(disableHeaderNamesNormalizing, false, "header names are passed as-is without normalization")
	config.Add(dialDualStack, "", "attempt to connect to both ipv4 and ipv6 addresses if set to true")
	config.Add(maxResponseBodySize, 52428800, "maximum response body size")
	config.Add(maxIdemponentCallAttempts, 5, "maximum number of attempts for idempotent calls")
}
