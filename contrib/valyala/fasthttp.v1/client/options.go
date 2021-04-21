package client

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	Name                          string
	NoDefaultUserAgentHeader      bool
	MaxConnsPerHost               int
	ReadBufferSize                int
	WriteBufferSize               int
	MaxConnWaitTimeout            time.Duration
	ReadTimeout                   time.Duration
	WriteTimeout                  time.Duration
	MaxIdleConnDuration           time.Duration
	DisableHeaderNamesNormalizing bool
	DialDualStack                 bool
	MaxResponseBodySize           int
	MaxIdemponentCallAttempts     int
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Name(value string) optionsBuilder {
	return builder.Set(b, "Name", value).(optionsBuilder)
}

func (b optionsBuilder) NoDefaultUserAgentHeader(value bool) optionsBuilder {
	return builder.Set(b, "NoDefaultUserAgentHeader", value).(optionsBuilder)
}

func (b optionsBuilder) MaxResponseBodySize(value int) optionsBuilder {
	return builder.Set(b, "MaxResponseBodySize", value).(optionsBuilder)
}

func (b optionsBuilder) MaxIdemponentCallAttempts(value int) optionsBuilder {
	return builder.Set(b, "MaxIdemponentCallAttempts", value).(optionsBuilder)
}

func (b optionsBuilder) MaxConnsPerHost(value int) optionsBuilder {
	return builder.Set(b, "MaxConnsPerHost", value).(optionsBuilder)
}

func (b optionsBuilder) WriteBufferSize(value int) optionsBuilder {
	return builder.Set(b, "WriteBufferSize", value).(optionsBuilder)
}

func (b optionsBuilder) MaxConnWaitTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxConnWaitTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) ReadTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "ReadTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) WriteTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "WriteTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) MaxIdleConnDuration(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxIdleConnDuration", value).(optionsBuilder)
}

func (b optionsBuilder) DisableHeaderNamesNormalizing(value bool) optionsBuilder {
	return builder.Set(b, "DisableHeaderNamesNormalizing", value).(optionsBuilder)
}

func (b optionsBuilder) DialDualStack(value bool) optionsBuilder {
	return builder.Set(b, "DialDualStack", value).(optionsBuilder)
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
