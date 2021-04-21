package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/b2wdigital/goignite/v2/core/errors"
	"github.com/b2wdigital/goignite/v2/core/log"
)

// Client knows how to publish on sns
type Client interface {
	Publish(ctx context.Context, input *sns.PublishInput) error
}

// Client holds client and resource name
type client struct {
	client *sns.Client
}

// NewClient returns a initialized client
func NewClient(c *sns.Client) Client {
	return &client{c}
}

// Publish publish message on sns
func (c *client) Publish(ctx context.Context, input *sns.PublishInput) error {

	logger := log.FromContext(ctx).
		WithTypeOf(*c).
		WithField("subject", input.Subject)

	logger.Tracef("sending message to sns with timeout")

	response, err := c.client.Publish(ctx, input)
	if err != nil {
		return errors.Wrap(err, errors.New("error publishing message on sns"))
	}

	logger.
		WithField("message_id", *response.MessageId).
		Debug("message sent to sns")

	return nil
}
