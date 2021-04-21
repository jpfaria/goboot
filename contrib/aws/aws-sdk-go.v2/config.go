package aws

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root                  = "gi.aws"
	key                   = root + ".accessKeyId"
	secret                = root + ".secretAccessKey"
	region                = root + ".defaultRegion"
	accountNumber         = root + ".defaultAccountNumber"
	customEndpoint        = root + ".customEndpoint"
	retryerRoot           = root + ".retryer"
	retryerMaxAttempts    = retryerRoot + ".maxAttempts"
	retryerHasRateLimit   = retryerRoot + ".hasRateLimit"
	ExtRoot               = root + ".ext"
	httpClientRoot        = root + ".httpClient"
	maxIdleConnPerHost    = httpClientRoot + ".maxIdleConnPerHost"
	maxIdleConn           = httpClientRoot + ".maxIdleConn"
	maxConnsPerHost       = httpClientRoot + ".maxConnsPerHost"
	idleConnTimeout       = httpClientRoot + ".idleConnTimeout"
	disableKeepAlives     = httpClientRoot + ".disableKeepAlives"
	disableCompression    = httpClientRoot + ".disableCompression"
	forceHTTP2            = httpClientRoot + ".forceHTTP2"
	tlsHandshakeTimeout   = httpClientRoot + ".TLSHandshakeTimeout"
	timeout               = httpClientRoot + ".timeout"
	dialTimeout           = httpClientRoot + ".dialTimeout"
	keepAlive             = httpClientRoot + ".keepAlive"
	expectContinueTimeout = httpClientRoot + ".expectContinueTimeout"
	dualStack             = httpClientRoot + ".dualStack"
)

func init() {
	config.Add(key, "", "defines the aws key id", config.WithHide())
	config.Add(secret, "", "defines the aws secret key", config.WithHide())
	config.Add(region, "", "defines the aws region")
	config.Add(accountNumber, "", "defines the aws account number")
	config.Add(customEndpoint, false, "defines if should point to localhost")
	config.Add(retryerMaxAttempts, 5, "defines max attempts for rate limit")
	config.Add(retryerHasRateLimit, true, "defines if retryer has rate limit")
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
