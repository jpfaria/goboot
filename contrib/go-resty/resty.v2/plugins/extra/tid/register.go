package tid

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/info"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling tid middleware in resty")

	client.OnBeforeRequest(tid)

	logger.Debug("tid middleware successfully enabled in resty")

	return nil
}

func tid(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	tidValue, ok := ctx.Value("x-tid").(string)
	if !ok {
		tidValue = info.AppName + "-" + uuid.NewV4().String()
	}

	request.SetHeader("X-TID", tidValue)

	return nil
}
