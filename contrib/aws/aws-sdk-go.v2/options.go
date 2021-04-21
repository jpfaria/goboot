package aws

import (
	"os"
	"time"

	"github.com/b2wdigital/goignite/v2/contrib/net/http/client"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
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

type optionsBuilder builder.Builder

func (b optionsBuilder) AccessKeyId(value string) optionsBuilder {
	return builder.Set(b, "AccessKeyId", value).(optionsBuilder)
}

func (b optionsBuilder) SecretAccessKey(value string) optionsBuilder {
	return builder.Set(b, "SecretAccessKey", value).(optionsBuilder)
}

func (b optionsBuilder) DefaultRegion(value string) optionsBuilder {
	return builder.Set(b, "DefaultRegion", value).(optionsBuilder)
}

func (b optionsBuilder) SessionToken(value string) optionsBuilder {
	return builder.Set(b, "SessionToken", value).(optionsBuilder)
}

func (b optionsBuilder) MaxAttempts(value int) optionsBuilder {
	return builder.Set(b, "MaxAttempts", value).(optionsBuilder)
}

func (b optionsBuilder) HasRateLimit(value bool) optionsBuilder {
	return builder.Set(b, "HasRateLimit", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func NewOptionsWithPath(path string) (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(path, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptions() (*Options, error) {
	o, err := NewOptionsWithPath(root)
	if err != nil {
		return nil, err
	}

	if v := os.Getenv("AWS_ACCESS_KEY_ID"); v != "" {
		o.AccessKeyId = v
	}

	if v := os.Getenv("AWS_SECRET_ACCESS_KEY"); v != "" {
		o.SecretAccessKey = v
	}

	if v := os.Getenv("AWS_DEFAULT_REGION"); v != "" {
		o.DefaultRegion = v
	}

	if v := os.Getenv("AWS_DEFAULT_ACCOUNT_NUMBER"); v != "" {
		o.DefaultAccountNumber = v
	}

	if v := os.Getenv("AWS_SESSION_TOKEN"); v != "" {
		o.SessionToken = v
	}

	return o, nil
}
