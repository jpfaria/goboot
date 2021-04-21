package redis

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	Sentinel           SentinelOptions
	Password           string
	MaxRetries         int
	MinRetryBackoff    time.Duration
	MaxRetryBackoff    time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
	Client             ClientOptions
	Cluster            ClusterOptions
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Password(value string) optionsBuilder {
	return builder.Set(b, "Password", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRetries(value int) optionsBuilder {
	return builder.Set(b, "MaxRetries", value).(optionsBuilder)
}

func (b optionsBuilder) MinRetryBackoff(value time.Duration) optionsBuilder {
	return builder.Set(b, "MinRetryBackoff", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRetryBackoff(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxRetryBackoff", value).(optionsBuilder)
}

func (b optionsBuilder) ReadTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "ReadTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) DialTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "DialTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) WriteTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "WriteTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) PoolSize(value int) optionsBuilder {
	return builder.Set(b, "PoolSize", value).(optionsBuilder)
}

func (b optionsBuilder) MinIdleConns(value int) optionsBuilder {
	return builder.Set(b, "MinIdleConns", value).(optionsBuilder)
}

func (b optionsBuilder) MaxConnAge(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxConnAge", value).(optionsBuilder)
}

func (b optionsBuilder) PoolTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "PoolTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) IdleTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "IdleTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) IdleCheckFrequency(value time.Duration) optionsBuilder {
	return builder.Set(b, "IdleCheckFrequency", value).(optionsBuilder)
}

func (b optionsBuilder) Client(value ClientOptions) optionsBuilder {
	return builder.Set(b, "Client", value).(optionsBuilder)
}

func (b optionsBuilder) Sentinel(value SentinelOptions) optionsBuilder {
	return builder.Set(b, "Sentinel", value).(optionsBuilder)
}

func (b optionsBuilder) Cluster(value ClusterOptions) optionsBuilder {
	return builder.Set(b, "Cluster", value).(optionsBuilder)
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
