package resty

import (
	"time"

	"github.com/lann/builder"
)

type OptionsTransport struct {
	DisableCompression    bool
	DisableKeepAlives     bool
	MaxIdleConnsPerHost   int
	ResponseHeaderTimeout time.Duration
	ForceAttemptHTTP2     bool `config:"forceAttemptHTTP2"`
	MaxIdleConns          int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	TLSHandshakeTimeout   time.Duration
	ExpectContinueTimeout time.Duration
}

type optionsTransportBuilder builder.Builder

func (b optionsTransportBuilder) DisableCompression(value bool) optionsTransportBuilder {
	return builder.Set(b, "DisableCompression", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) DisableKeepAlives(value bool) optionsTransportBuilder {
	return builder.Set(b, "DisableKeepAlives", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) MaxIdleConnsPerHost(value int) optionsTransportBuilder {
	return builder.Set(b, "MaxIdleConnsPerHost", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) ResponseHeaderTimeout(value time.Duration) optionsTransportBuilder {
	return builder.Set(b, "ResponseHeaderTimeout", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) ForceAttemptHTTP2(value bool) optionsTransportBuilder {
	return builder.Set(b, "ForceAttemptHTTP2", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) MaxIdleConns(value int) optionsTransportBuilder {
	return builder.Set(b, "MaxIdleConns", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) MaxConnsPerHost(value int) optionsTransportBuilder {
	return builder.Set(b, "MaxConnsPerHost", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) IdleConnTimeout(value time.Duration) optionsTransportBuilder {
	return builder.Set(b, "IdleConnTimeout", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) TLSHandshakeTimeout(value time.Duration) optionsTransportBuilder {
	return builder.Set(b, "TLSHandshakeTimeout", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) ExpectContinueTimeout(value time.Duration) optionsTransportBuilder {
	return builder.Set(b, "ExpectContinueTimeout", value).(optionsTransportBuilder)
}

func (b optionsTransportBuilder) Build() OptionsTransport {
	return builder.GetStruct(b).(OptionsTransport)
}

var OptionsTransportBuilder = builder.Register(optionsTransportBuilder{}, Options{}).(optionsTransportBuilder)
