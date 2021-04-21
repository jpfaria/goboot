package client

import (
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	TLS                bool `config:"tls"`
	Gzip               bool
	CertFile           string
	KeyFile            string
	CAFile             string `config:"caFile"`
	Host               string
	HostOverwrite      string
	Port               int
	InsecureSkipVerify bool
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Gzip(enabled bool) optionsBuilder {
	return builder.Set(b, "Gzip", enabled).(optionsBuilder)
}

func (b optionsBuilder) Tls(enabled bool) optionsBuilder {
	return builder.Set(b, "Tls", enabled).(optionsBuilder)
}

func (b optionsBuilder) InsecureSkipVerify(skip bool) optionsBuilder {
	return builder.Set(b, "InsecureSkipVerify", skip).(optionsBuilder)
}

func (b optionsBuilder) CertFile(file string) optionsBuilder {
	return builder.Set(b, "CertFile", file).(optionsBuilder)
}

func (b optionsBuilder) KeyFile(file string) optionsBuilder {
	return builder.Set(b, "KeyFile", file).(optionsBuilder)
}

func (b optionsBuilder) CAFile(file string) optionsBuilder {
	return builder.Set(b, "CAFile", file).(optionsBuilder)
}

func (b optionsBuilder) Host(host string) optionsBuilder {
	return builder.Set(b, "Host", host).(optionsBuilder)
}

func (b optionsBuilder) HostOverwrite(hostOverwrite string) optionsBuilder {
	return builder.Set(b, "HostOverwrite", hostOverwrite).(optionsBuilder)
}

func (b optionsBuilder) Port(port int) optionsBuilder {

	if port == 0 {
		port = 9090
	}

	return builder.Set(b, "Port", port).(optionsBuilder)
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
