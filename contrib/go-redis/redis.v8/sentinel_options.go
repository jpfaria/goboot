package redis

import "github.com/lann/builder"

type SentinelOptions struct {
	MasterName string `config:"masterName"`
	Addrs      []string
	Password   string
}

type sentinelOptionsBuilder builder.Builder

func (b sentinelOptionsBuilder) MasterName(value string) sentinelOptionsBuilder {
	return builder.Set(b, "MasterName", value).(sentinelOptionsBuilder)
}

func (b sentinelOptionsBuilder) Addrs(value []string) sentinelOptionsBuilder {
	return builder.Set(b, "Addrs", value).(sentinelOptionsBuilder)
}

func (b sentinelOptionsBuilder) Password(value string) sentinelOptionsBuilder {
	return builder.Set(b, "Password", value).(sentinelOptionsBuilder)
}

func (b sentinelOptionsBuilder) Build() ClientOptions {
	return builder.GetStruct(b).(ClientOptions)
}

var SentinelOptionsBuilder = builder.Register(sentinelOptionsBuilder{}, SentinelOptions{}).(sentinelOptionsBuilder)
