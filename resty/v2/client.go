package giresty

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-resty/resty/v2"
)

const (
	TopicClient = "topic:resty:client"
)

func NewClient(ctx context.Context, options *Options) *resty.Client {

	l := gilog.FromContext(ctx)

	l.Infof("creating resty client")

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:   giconfig.Duration(connectionTimeout),
		KeepAlive: giconfig.Duration(keepAlive),
		DualStack: true,
	}

	if options.ConnectionTimeout > 0 {
		dialer.Timeout = options.ConnectionTimeout
	}

	if options.KeepAlive > 0 {
		dialer.KeepAlive = options.KeepAlive
	}

	transport := &http.Transport{
		DisableCompression:    giconfig.Bool(transportDisableCompression),
		DisableKeepAlives:     giconfig.Bool(transportDisableKeepAlives),
		MaxIdleConnsPerHost:   giconfig.Int(transportMaxConnsPerHost),
		ResponseHeaderTimeout: giconfig.Duration(transportResponseHeaderTimeout),
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     giconfig.Bool(transportForceAttemptHTTP2),
		MaxIdleConns:          giconfig.Int(transportMaxIdleConns),
		MaxConnsPerHost:       giconfig.Int(transportMaxConnsPerHost),
		IdleConnTimeout:       giconfig.Duration(transportIdleConnTimeout),
		TLSHandshakeTimeout:   giconfig.Duration(transportTLSHandshakeTimeout),
		ExpectContinueTimeout: giconfig.Duration(transportExpectContinueTimeout),
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
		SetTimeout(giconfig.Duration(requestTimeout)).
		SetRetryCount(giconfig.Int(retryCount)).
		SetRetryWaitTime(giconfig.Duration(retryWaitTime)).
		SetRetryMaxWaitTime(giconfig.Duration(retryMaxWaitTime)).
		SetDebug(giconfig.Bool(debug)).
		SetHostURL(options.Host).
		SetCloseConnection(giconfig.Bool(closeConnection)).
		AddRetryCondition(statusCodeRetryCondition)

	addTimeoutRetryCondition(client, options)

	if options.Debug || giconfig.Bool(debug) {
		client.OnBeforeRequest(logBeforeResponse)
		client.OnAfterResponse(logAfterResponse)
		client.SetDebug(true)
	}

	if options.RequestTimeout > 0 {
		client.SetTimeout(options.RequestTimeout)
	}

	if options.Retry != nil {

		if options.Retry.Count > 0 {
			client.SetRetryCount(options.Retry.Count)
		}

		if options.Retry.WaitTime > 0 {
			client.SetRetryWaitTime(options.Retry.WaitTime)
		}

		if options.Retry.MaxWaitTime > 0 {
			client.SetRetryMaxWaitTime(options.Retry.MaxWaitTime)
		}
	}

	gieventbus.Publish(TopicClient, client)

	return client
}

func addTimeoutRetryCondition(client *resty.Client, options *Options) {

	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {

			var timeout time.Duration

			if options.RequestTimeout > 0 {
				timeout = options.RequestTimeout
			} else {
				timeout = giconfig.Duration(requestTimeout)
			}

			if r.Time() > timeout {
				return true
			}

			return false
		})
}

func statusCodeRetryCondition(r *resty.Response, err error) bool {
	switch statusCode := r.StatusCode(); statusCode {

	case http.StatusTooManyRequests:
		return true
	case http.StatusInternalServerError:
		return true
	case http.StatusGatewayTimeout:
		return true
	case http.StatusServiceUnavailable:
		return true
	default:
		return false
	}
}

func logBeforeResponse(client *resty.Client, request *resty.Request) error {

	l := gilog.FromContext(request.Context())

	requestHeaders, _ := json.Marshal(request.Header)

	requestBody, _ := json.Marshal(request.Body)

	l = l.
		WithFields(
			gilog.Fields{
				"rest_request_body":    string(requestBody),
				"rest_request_url":     request.URL,
				"rest_request_headers": string(requestHeaders),
				"rest_request_method":  request.Method,
			})

	l.Debugf("rest request processing")

	return nil
}

func logAfterResponse(client *resty.Client, response *resty.Response) error {

	l := gilog.FromContext(response.Request.Context())

	responseHeaders, _ := json.Marshal(response.Header())

	statusCode := response.StatusCode()

	l = l.WithFields(
		gilog.Fields{
			"rest_response_body":        string(response.Body()),
			"rest_response_headers":     string(responseHeaders),
			"rest_response_time":        response.Time().Seconds() * float64(time.Second/time.Millisecond),
			"rest_response_status_code": statusCode,
		})

	if statusCode > 500 {
		l.Errorf("rest request processed with error")
	} else if statusCode > 400 {
		l.Warnf("rest request processed with warning")
	} else {
		l.Debugf("successful rest request processed")
	}

	return nil
}
