package main

import (
	"context"

	akinesis "github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2"
	"github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2/client/kinesis"
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
	sqsClient := akinesis.NewFromConfig(awsConfig)
	client := kinesis.NewClient(sqsClient)

	input := &akinesis.PutRecordInput{
		Data:                      nil,
		PartitionKey:              nil,
		StreamName:                nil,
		ExplicitHashKey:           nil,
		SequenceNumberForOrdering: nil,
	}

	// publish
	err := client.Publish(ctx, input)
	if err != nil {
		logger.Fatalf(err.Error())
	}

}
