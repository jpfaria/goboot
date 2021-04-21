package elasticsearch

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	Addresses             []string
	Username              string
	Password              string
	CloudID               string `config:"cloudID"`
	APIKey                string `config:"APIKey"`
	CACert                string `config:"CACert"`
	RetryOnStatus         []int
	DisableRetry          bool
	EnableRetryOnTimeout  bool
	MaxRetries            int
	DiscoverNodesOnStart  bool
	DiscoverNodesInterval time.Duration
	EnableMetrics         bool
	EnableDebugLogger     bool
	RetryBackoff          time.Duration
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Addresses(value []string) optionsBuilder {
	return builder.Set(b, "Addresses", value).(optionsBuilder)
}

func (b optionsBuilder) Username(value string) optionsBuilder {
	return builder.Set(b, "Username", value).(optionsBuilder)
}

func (b optionsBuilder) Password(value string) optionsBuilder {
	return builder.Set(b, "Password", value).(optionsBuilder)
}

func (b optionsBuilder) CloudID(value string) optionsBuilder {
	return builder.Set(b, "CloudID", value).(optionsBuilder)
}

func (b optionsBuilder) APIKey(value string) optionsBuilder {
	return builder.Set(b, "APIKey", value).(optionsBuilder)
}

func (b optionsBuilder) CACert(value string) optionsBuilder {
	return builder.Set(b, "CACert", value).(optionsBuilder)
}

func (b optionsBuilder) RetryOnStatus(value []int) optionsBuilder {
	return builder.Set(b, "RetryOnStatus", value).(optionsBuilder)
}

func (b optionsBuilder) DisableRetry(value bool) optionsBuilder {
	return builder.Set(b, "DisableRetry", value).(optionsBuilder)
}

func (b optionsBuilder) EnableRetryOnTimeout(value bool) optionsBuilder {
	return builder.Set(b, "EnableRetryOnTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRetries(value int) optionsBuilder {
	return builder.Set(b, "MaxRetries", value).(optionsBuilder)
}

func (b optionsBuilder) DiscoverNodesOnStart(value bool) optionsBuilder {
	return builder.Set(b, "DiscoverNodesOnStart", value).(optionsBuilder)
}

func (b optionsBuilder) DiscoverNodesInterval(value time.Duration) optionsBuilder {
	return builder.Set(b, "DiscoverNodesInterval", value).(optionsBuilder)
}

func (b optionsBuilder) EnableMetrics(value bool) optionsBuilder {
	return builder.Set(b, "EnableMetrics", value).(optionsBuilder)
}

func (b optionsBuilder) EnableDebugLogger(value bool) optionsBuilder {
	return builder.Set(b, "EnableDebugLogger", value).(optionsBuilder)
}

func (b optionsBuilder) RetryBackoff(value time.Duration) optionsBuilder {
	return builder.Set(b, "RetryBackoff", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
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

	return opts, nil
}
