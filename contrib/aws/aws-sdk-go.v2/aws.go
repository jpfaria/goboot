package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	c "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/b2wdigital/goignite/v2/contrib/net/http/client"
	"github.com/b2wdigital/goignite/v2/core/log"
)

type Ext func(context.Context, *aws.Config) error

func NewConfig(ctx context.Context, options *Options, exts ...Ext) aws.Config {

	logger := log.FromContext(ctx)

	cfg, err := c.LoadDefaultConfig(ctx)
	if err != nil {
		logger.Panicf("unable to load AWS SDK config, %s", err.Error())
		return aws.Config{}
	}

	cfg.Region = options.DefaultRegion
	cfg.Credentials = credentials.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)

	httpClient := client.NewClient(ctx, &options.HttpClient)

	cfg.Retryer = retryerConfig(options)
	cfg.HTTPClient = httpClient

	for _, ext := range exts {
		if err := ext(ctx, &cfg); err != nil {
			panic(err)
		}
	}

	return cfg
}

func retryerConfig(options *Options) func() aws.Retryer {

	return func() aws.Retryer {

		return retry.NewStandard(func(o *retry.StandardOptions) {

			o.MaxAttempts = options.MaxAttempts

			if !options.HasRateLimit {
				o.RateLimiter = noRateLimit{}
			}

		})
	}
}

type noRateLimit struct{}

func (noRateLimit) AddTokens(uint) error { return nil }

func (noRateLimit) GetToken(context.Context, uint) (func() error, error) { return nil, nil }

func NewDefaultConfig(ctx context.Context, exts ...Ext) aws.Config {

	o, err := DefaultOptions()
	if err != nil {
		panic(err)
	}

	return NewConfig(ctx, o, exts...)
}
