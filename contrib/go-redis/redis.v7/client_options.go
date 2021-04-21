package redis

import (
	"github.com/lann/builder"
)

type ClientOptions struct {
	Addr    string
	Network string
	DB      int `config:"db"`
}

type clientOptionsBuilder builder.Builder

func (b clientOptionsBuilder) Addr(value string) clientOptionsBuilder {
	return builder.Set(b, "addr", value).(clientOptionsBuilder)
}

func (b clientOptionsBuilder) Network(value string) clientOptionsBuilder {
	return builder.Set(b, "network", value).(clientOptionsBuilder)
}

func (b clientOptionsBuilder) DB(value int) clientOptionsBuilder {
	return builder.Set(b, "db", value).(clientOptionsBuilder)
}

func (b clientOptionsBuilder) Build() ClientOptions {
	return builder.GetStruct(b).(ClientOptions)
}

var ClientOptionsBuilder = builder.Register(clientOptionsBuilder{}, ClientOptions{}).(clientOptionsBuilder)
