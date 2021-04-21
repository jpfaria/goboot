package mongo

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	Uri  string
	Auth *options.Credential
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Uri(value string) optionsBuilder {
	return builder.Set(b, "Uri", value).(optionsBuilder)
}

func (b optionsBuilder) Auth(value *options.Credential) optionsBuilder {
	return builder.Set(b, "Auth", value).(optionsBuilder)
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
