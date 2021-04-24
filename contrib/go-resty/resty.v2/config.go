package resty

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root                           = "gi.resty"
	debug                          = root + ".debug"
	closeConnection                = root + ".closeConnection"
	connectionTimeout              = root + ".connectionTimeout"
	keepAlive                      = root + ".keepAlive"
	fallbackDelay                  = root + ".fallbackDelay"
	requestTimeout                 = root + ".requestTimeout"
	transportDisableCompression    = root + ".transport.disableCompression"
	transportDisableKeepAlives     = root + ".transport.disableKeepAlives"
	transportMaxIdleConnsPerHost   = root + ".transport.maxIdleConnsPerHost"
	transportResponseHeaderTimeout = root + ".transport.responseHeaderTimeout"
	transportForceAttemptHTTP2     = root + ".transport.forceAttemptHTTP2"
	transportMaxIdleConns          = root + ".transport.maxIdleConns"
	transportMaxConnsPerHost       = root + ".transport.maxConnsPerHost"
	transportIdleConnTimeout       = root + ".transport.idleConnTimeout"
	transportTLSHandshakeTimeout   = root + ".transport.TLSHandshakeTimeout"
	transportExpectContinueTimeout = root + ".transport.expectContinueTimeout"
	PluginsRoot                    = root + ".plugins"
)

func init() {

	config.Add(debug, false, "defines global debug request")
	config.Add(closeConnection, false, "defines global http close connection")
	config.Add(connectionTimeout, 3*time.Minute, "defines global http connection timeout")
	config.Add(keepAlive, 30*time.Second, "defines global http keepalive")
	config.Add(fallbackDelay, 300*time.Millisecond, "defines global fallbackDelay")
	config.Add(requestTimeout, 30*time.Second, "defines global http request timeout")
	config.Add(transportDisableCompression, false, "enabled/disable transport compression")
	config.Add(transportDisableKeepAlives, false, "enabled/disable transport keep alives")
	config.Add(transportMaxIdleConnsPerHost, 2, "define transport max idle conns per host")
	config.Add(transportResponseHeaderTimeout, 2*time.Second, "define transport response header timeout")
	config.Add(transportForceAttemptHTTP2, true, "define transport force attempt http2")
	config.Add(transportMaxIdleConns, 100, "define transport max idle conns")
	config.Add(transportMaxConnsPerHost, 100, "define transport max conns per host")
	config.Add(transportIdleConnTimeout, 90*time.Second, "define transport idle conn timeout")
	config.Add(transportTLSHandshakeTimeout, 10*time.Second, "define transport TLS handshake timeout")
	config.Add(transportExpectContinueTimeout, 1*time.Second, "define transport expect continue timeout")
}
