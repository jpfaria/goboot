package client

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root                  = "gi.http.client"
	maxIdleConnPerHost    = root + ".maxIdleConnPerHost"
	maxIdleConn           = root + ".maxIdleConn"
	maxConnsPerHost       = root + ".maxConnsPerHost"
	idleConnTimeout       = root + ".idleConnTimeout"
	disableKeepAlives     = root + ".disableKeepAlives"
	disableCompression    = root + ".disableCompression"
	forceHTTP2            = root + ".forceHTTP2"
	tlsHandshakeTimeout   = root + ".TLSHandshakeTimeout"
	timeout               = root + ".timeout"
	dialTimeout           = root + ".dialTimeout"
	keepAlive             = root + ".keepAlive"
	expectContinueTimeout = root + ".expectContinueTimeout"
	dualStack             = root + ".dualStack"
)

func init() {

	config.Add(maxIdleConnPerHost, 1, "http max idle connections per host")
	config.Add(maxIdleConn, 100, "http max idle connections")
	config.Add(maxConnsPerHost, 20, "http max connections per host")
	config.Add(idleConnTimeout, 90*time.Second, "http idle connections timeout")
	config.Add(disableKeepAlives, true, "http disable keep alives")
	config.Add(disableCompression, false, "http disable keep alives")
	config.Add(forceHTTP2, true, "http force http2")
	config.Add(tlsHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	config.Add(timeout, 30*time.Second, "timeout")
	config.Add(dialTimeout, 5*time.Second, "dial timeout")
	config.Add(keepAlive, 15*time.Second, "keep alive")
	config.Add(expectContinueTimeout, 1*time.Second, "expect continue timeout")
	config.Add(dualStack, true, "dual stack")

}
