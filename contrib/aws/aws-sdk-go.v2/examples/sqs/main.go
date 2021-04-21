package main

import (
	"context"

	asqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2"
	"github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2/client/sqs"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/log"
)

const Bucket = "aws.s3.bucket"

func init() {
	config.Add(Bucket, "example", "s3 example bucket")
}

func main() {

	config.Load()

	// create background context
	ctx := context.Background()

	// start logrus
	// zap.NewLogger()
	logrus.NewLogger()

	// get logrus instance from context
	logger := log.FromContext(ctx)

	// create default aws config
	awsConfig := aws.NewConfig(ctx)

	// create sns client
	sqsClient := asqs.NewFromConfig(awsConfig)
	client := sqs.NewClient(sqsClient)

	input := &asqs.SendMessageInput{
		MessageBody:             nil,
		QueueUrl:                nil,
		DelaySeconds:            0,
		MessageAttributes:       nil,
		MessageDeduplicationId:  nil,
		MessageGroupId:          nil,
		MessageSystemAttributes: nil,
	}

	// publish
	err := client.Publish(ctx, input)
	if err != nil {
		logger.Fatalf(err.Error())
	}

}
