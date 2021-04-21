package retry

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-resty/resty/v2"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("configuring retry in resty")

	client.
		SetRetryCount(config.Int(count)).
		SetRetryWaitTime(config.Duration(waitTime)).
		SetRetryMaxWaitTime(config.Duration(maxWaitTime)).
		AddRetryCondition(statusCodeRetryCondition).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {

				timeout := client.GetClient().Timeout

				if r.Time() > timeout {
					return true
				}

				return false
			})

	addTimeoutRetryCondition(client)

	logger.Debug("retry successfully configured in resty")

	return nil
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

func addTimeoutRetryCondition(client *resty.Client) {

	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {

			timeout := client.GetClient().Timeout

			if r.Time() > timeout {
				return true
			}

			return false
		})
}
