package health

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	Name        string
	Host        string
	Endpoint    string
	Enabled     bool
	Description string
	Required    bool
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Name(value string) optionsBuilder {
	return builder.Set(b, "Name", value).(optionsBuilder)
}

func (b optionsBuilder) Host(value string) optionsBuilder {
	return builder.Set(b, "Host", value).(optionsBuilder)
}

func (b optionsBuilder) Endpoint(value string) optionsBuilder {
	return builder.Set(b, "Endpoint", value).(optionsBuilder)
}

func (b optionsBuilder) Enabled(value bool) optionsBuilder {
	return builder.Set(b, "Enabled", value).(optionsBuilder)
}

func (b optionsBuilder) Description(value string) optionsBuilder {
	return builder.Set(b, "Description", value).(optionsBuilder)
}

func (b optionsBuilder) Required(value bool) optionsBuilder {
	return builder.Set(b, "Required", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
