package retry

import (
	"context"
	"net/http"
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
	r "github.com/go-resty/resty/v2"
)

func Register(ctx context.Context, client *r.Client) error {

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
		AddRetryCondition(addTimeoutRetryCondition(client.GetClient().Timeout))

	logger.Debug("retry successfully configured in resty")

	return nil
}

func statusCodeRetryCondition(r *r.Response, err error) bool {
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

func addTimeoutRetryCondition(timeout time.Duration) func(r *r.Response, err error) bool {

	return func(resp *r.Response, err error) bool {

		if resp.Time() > timeout {
			return true
		}

		return false
	}
}
