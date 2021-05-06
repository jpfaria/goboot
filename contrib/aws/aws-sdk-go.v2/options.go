package aws

import (
	"os"
	"time"

	"github.com/b2wdigital/goignite/v2/contrib/net/http/client"
	"github.com/b2wdigital/goignite/v2/core/config"
)

type Options struct {
	AccessKeyId                 string
	SecretAccessKey             string
	DefaultRegion               string
	SessionToken                string
	DefaultAccountNumber        string
	MaxAttempts                 int
	HasRateLimit                bool
	MaxConnsPerHost             int
	MaxIdleConns                int
	MaxIdleConnsPerHost         int
	TimeoutMillis               time.Duration
	KeepAliveMillis             time.Duration
	IdleConnTimeoutMillis       time.Duration
	ResponseHeaderTimeoutMillis time.Duration
	HttpClient                  client.Options
}

func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	loadEnv(opts)

	return opts, nil
}

func NewOptions() (*Options, error) {
	opts := &Options{}

	err := config.UnmarshalWithPath(root, opts)
	if err != nil {
		return nil, err
	}

	loadEnv(opts)

	return opts, nil
}

func loadEnv(opts *Options) {

	if v := os.Getenv("AWS_ACCESS_KEY_ID"); v != "" {
		opts.AccessKeyId = v
	}

	if v := os.Getenv("AWS_SECRET_ACCESS_KEY"); v != "" {
		opts.SecretAccessKey = v
	}

	if v := os.Getenv("AWS_DEFAULT_REGION"); v != "" {
		opts.DefaultRegion = v
	}

	if v := os.Getenv("AWS_DEFAULT_ACCOUNT_NUMBER"); v != "" {
		opts.DefaultAccountNumber = v
	}

	if v := os.Getenv("AWS_SESSION_TOKEN"); v != "" {
		opts.SessionToken = v
	}

}
