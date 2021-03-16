package girestyhealth

import (
	"context"
	"strconv"
	"strings"

	gierrors "github.com/b2wdigital/goignite/v2/errors"
	"github.com/go-resty/resty/v2"
)

type Checker struct {
	client  *resty.Client
	options *Options
}

func (c *Checker) Check(ctx context.Context) (err error) {

	request := c.client.R().EnableTrace()

	var response *resty.Response

	response, err = request.Get(strings.Join([]string{c.options.Host, c.options.Endpoint}, ""))
	if err != nil {
		return gierrors.Internalf(err.Error())
	}

	if response.IsError() {
		return gierrors.New(strconv.Itoa(response.StatusCode()))
	}

	return err
}

func NewChecker(client *resty.Client, options *Options) *Checker {
	return &Checker{client: client, options: options}
}
