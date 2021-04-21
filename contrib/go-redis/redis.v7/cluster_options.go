package redis

import (
	"github.com/lann/builder"
)

type ClusterOptions struct {
	Addrs          []string
	MaxRedirects   int  `config:"maxredirects"`
	ReadOnly       bool `config:"readonly"`
	RouteByLatency bool `config:"routebylatency"`
	RouteRandomly  bool `config:"routerandomly"`
}

type clusterOptionsBuilder builder.Builder

func (b clusterOptionsBuilder) Addrs(value []string) clusterOptionsBuilder {
	return builder.Set(b, "addrs", value).(clusterOptionsBuilder)
}

func (b clusterOptionsBuilder) MaxRedirects(value int) clusterOptionsBuilder {
	return builder.Set(b, "maxRedirects", value).(clusterOptionsBuilder)
}

func (b clusterOptionsBuilder) ReadOnly(value bool) clusterOptionsBuilder {
	return builder.Set(b, "readOnly", value).(clusterOptionsBuilder)
}

func (b clusterOptionsBuilder) RouteByLatency(value bool) clusterOptionsBuilder {
	return builder.Set(b, "routeByLatency", value).(clusterOptionsBuilder)
}

func (b clusterOptionsBuilder) Build() ClusterOptions {
	return builder.GetStruct(b).(ClusterOptions)
}

var ClusterOptionsBuilder = builder.Register(clusterOptionsBuilder{}, ClusterOptions{}).(clusterOptionsBuilder)
