package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/b2wdigital/goignite/v2/core/errors"
	"github.com/b2wdigital/goignite/v2/core/log"
)

// Client knows how to publish on sqs
type Client interface {
	Publish(ctx context.Context, input *sqs.SendMessageInput) error
}

// Client holds client and resource name
type client struct {
	client *sqs.Client
}

// NewClient returns a initialized client
func NewClient(c *sqs.Client) Client {
	return &client{c}
}

// Publish publish message on sns
func (c *client) Publish(ctx context.Context, input *sqs.SendMessageInput) error {

	logger := log.FromContext(ctx).
		WithTypeOf(*c).
		WithField("subject", input.QueueUrl)

	logger.Tracef("sending message to sqs")

	response, err := c.client.SendMessage(ctx, input)
	if err != nil {
		return errors.Wrap(err, errors.New("error sending message to sqs"))
	}

	logger.
		WithField("message_id", *response.MessageId).
		Debug("message sent to sqs")

	return nil
}
