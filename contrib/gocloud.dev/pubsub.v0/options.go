package pubsub

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

// Options ..
type Options struct {
	Resource string
	Type     string
	Region   string
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Resource(resource string) optionsBuilder {
	return builder.Set(b, "Resource", resource).(optionsBuilder)
}

func (b optionsBuilder) Type(tp string) optionsBuilder {
	return builder.Set(b, "Type", tp).(optionsBuilder)
}

func (b optionsBuilder) Region(region string) optionsBuilder {
	return builder.Set(b, "Region", region).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

// OptionsBuilder ..
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
