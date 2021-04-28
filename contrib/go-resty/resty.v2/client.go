package resty

import (
	"context"
	"net"
	"net/http"

	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-resty/resty/v2"
)

type Ext func(context.Context, *resty.Client) error

func NewClientWithOptions(ctx context.Context, options *Options, exts ...Ext) *resty.Client {

	logger := log.FromContext(ctx)

	logger.Tracef("creating resty client for host %s", options.Host)

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:       options.ConnectionTimeout,
		FallbackDelay: options.FallbackDelay,
		KeepAlive:     options.KeepAlive,
	}

	transport := &http.Transport{
		DisableCompression:    options.Transport.DisableCompression,
		DisableKeepAlives:     options.Transport.DisableKeepAlives,
		MaxIdleConnsPerHost:   options.Transport.MaxIdleConnsPerHost,
		ResponseHeaderTimeout: options.Transport.ResponseHeaderTimeout,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     options.Transport.ForceAttemptHTTP2,
		MaxIdleConns:          options.Transport.MaxIdleConns,
		MaxConnsPerHost:       options.Transport.MaxConnsPerHost,
		IdleConnTimeout:       options.Transport.IdleConnTimeout,
		TLSHandshakeTimeout:   options.Transport.TLSHandshakeTimeout,
		ExpectContinueTimeout: options.Transport.ExpectContinueTimeout,
	}

	client.
		SetTransport(transport).
		SetTimeout(options.RequestTimeout).
		SetDebug(options.Debug).
		SetHostURL(options.Host).
		SetCloseConnection(options.CloseConnection)

	for _, ext := range exts {
		if err := ext(ctx, client); err != nil {
			panic(err)
		}
	}

	logger.Debugf("resty client created for host %s", options.Host)

	return client
}
