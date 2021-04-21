package godror

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	DataSourceName  string        `config:"datasourcename"`
	ConnMaxLifetime time.Duration `config:"connmaxlifetime"`
	MaxIdleConns    int           `config:"maxidleconns"`
	MaxOpenConns    int           `config:"maxopenconns"`
}

type optionsBuilder builder.Builder

func (b optionsBuilder) DataSourceName(value string) optionsBuilder {
	return builder.Set(b, "DataSourceName", value).(optionsBuilder)
}

func (b optionsBuilder) ConnMaxLifetime(value time.Duration) optionsBuilder {
	return builder.Set(b, "ConnMaxLifetime", value).(optionsBuilder)
}

func (b optionsBuilder) MaxIdleConns(value int) optionsBuilder {
	return builder.Set(b, "MaxIdleConns", value).(optionsBuilder)
}

func (b optionsBuilder) MaxOpenConns(value int) optionsBuilder {
	return builder.Set(b, "MaxOpenConns", value).(optionsBuilder)
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
