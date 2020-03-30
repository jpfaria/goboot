package jlaffaye

import (
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/lann/builder"
)

type Options struct {
	Addr     string
	User     string
	Password string
	Timeout  int
	Retry    int
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Addr(value string) optionsBuilder {
	return builder.Set(b, "Addr", value).(optionsBuilder)
}

func (b optionsBuilder) User(value string) optionsBuilder {
	return builder.Set(b, "User", value).(optionsBuilder)
}

func (b optionsBuilder) Password(value string) optionsBuilder {
	return builder.Set(b, "Password", value).(optionsBuilder)
}

func (b optionsBuilder) Timeout(value int) optionsBuilder {
	return builder.Set(b, "Timeout", value).(optionsBuilder)
}

func (b optionsBuilder) Retry(value int) optionsBuilder {
	return builder.Set(b, "Retry", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath("transport.client.jlaffaye", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
