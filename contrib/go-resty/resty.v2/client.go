package resty

import (
	"context"
	"net"
	"net/http"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-resty/resty/v2"
)

type Ext func(context.Context, *resty.Client) error

func NewClientWithOptions(ctx context.Context, options *Options, exts ...Ext) *resty.Client {

	logger := log.FromContext(ctx)

	logger.Tracef("creating resty client for host %s", options.Host)

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:       config.Duration(connectionTimeout),
		FallbackDelay: config.Duration(fallbackDelay),
		KeepAlive:     config.Duration(keepAlive),
	}

	if options.ConnectionTimeout > 0 {
		dialer.Timeout = options.ConnectionTimeout
	}

	if options.KeepAlive > 0 {
		dialer.KeepAlive = options.KeepAlive
	}

	transport := &http.Transport{
		DisableCompression:    config.Bool(transportDisableCompression),
		DisableKeepAlives:     config.Bool(transportDisableKeepAlives),
		MaxIdleConnsPerHost:   config.Int(transportMaxConnsPerHost),
		ResponseHeaderTimeout: config.Duration(transportResponseHeaderTimeout),
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     config.Bool(transportForceAttemptHTTP2),
		MaxIdleConns:          config.Int(transportMaxIdleConns),
		MaxConnsPerHost:       config.Int(transportMaxConnsPerHost),
		IdleConnTimeout:       config.Duration(transportIdleConnTimeout),
		TLSHandshakeTimeout:   config.Duration(transportTLSHandshakeTimeout),
		ExpectContinueTimeout: config.Duration(transportExpectContinueTimeout),
	}

	if options.Transport != nil {

		transport.DisableCompression = options.Transport.DisableCompression
		transport.DisableKeepAlives = options.Transport.DisableKeepAlives
		transport.ForceAttemptHTTP2 = options.Transport.ForceAttemptHTTP2

		if options.Transport.MaxIdleConnsPerHost > 0 {
			transport.MaxIdleConnsPerHost = options.Transport.MaxIdleConnsPerHost
		}

		if options.Transport.ResponseHeaderTimeout > 0 {
			transport.ResponseHeaderTimeout = options.Transport.ResponseHeaderTimeout
		}

		if options.Transport.MaxIdleConns > 0 {
			transport.MaxIdleConns = options.Transport.MaxIdleConns
		}

		if options.Transport.MaxConnsPerHost > 0 {
			transport.MaxConnsPerHost = options.Transport.MaxConnsPerHost
		}

		if options.Transport.IdleConnTimeout > 0 {
			transport.IdleConnTimeout = options.Transport.IdleConnTimeout
		}

		if options.Transport.TLSHandshakeTimeout > 0 {
			transport.TLSHandshakeTimeout = options.Transport.TLSHandshakeTimeout
		}

		if options.Transport.ExpectContinueTimeout > 0 {
			transport.ExpectContinueTimeout = options.Transport.ExpectContinueTimeout
		}
	}

	client.
		SetTransport(transport).
		SetTimeout(config.Duration(requestTimeout)).
		SetDebug(config.Bool(debug)).
		SetHostURL(options.Host).
		SetCloseConnection(config.Bool(closeConnection))

	if options.Debug || config.Bool(debug) {
		client.SetDebug(true)
	}

	if options.RequestTimeout > 0 {
		client.SetTimeout(options.RequestTimeout)
	}

	for _, ext := range exts {
		if err := ext(ctx, client); err != nil {
			panic(err)
		}
	}

	logger.Debugf("resty client created for host %s", options.Host)

	return client
}
