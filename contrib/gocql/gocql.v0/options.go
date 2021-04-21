package gocql

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/lann/builder"
)

type Options struct {
	Hosts                    []string
	Port                     int
	DC                       string `config:"dc"`
	Username                 string
	Password                 string
	CQLVersion               string `config:"CQLVersion"`
	ProtoVersion             int
	Timeout                  time.Duration
	ConnectTimeout           time.Duration
	Keyspace                 string
	NumConns                 int
	Consistency              string
	SocketKeepalive          time.Duration
	MaxPreparedStmts         int
	MaxRoutingKeyInfo        int
	PageSize                 int
	DefaultTimestamp         bool
	ReconnectInterval        time.Duration
	MaxWaitSchemaAgreement   time.Duration
	DisableInitialHostLookup bool
	WriteCoalesceWaitTime    time.Duration
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Hosts(value []string) optionsBuilder {
	return builder.Set(b, "Hosts", value).(optionsBuilder)
}

func (b optionsBuilder) Port(value int) optionsBuilder {
	return builder.Set(b, "Port", value).(optionsBuilder)
}

func (b optionsBuilder) DC(value string) optionsBuilder {
	return builder.Set(b, "DC", value).(optionsBuilder)
}

func (b optionsBuilder) Username(value int) optionsBuilder {
	return builder.Set(b, "Username", value).(optionsBuilder)
}

func (b optionsBuilder) Password(value int) optionsBuilder {
	return builder.Set(b, "Password", value).(optionsBuilder)
}

func (b optionsBuilder) CQLVersion(value string) optionsBuilder {
	return builder.Set(b, "CQLVersion", value).(optionsBuilder)
}

func (b optionsBuilder) ProtoVersion(value int) optionsBuilder {
	return builder.Set(b, "ProtoVersion", value).(optionsBuilder)
}

func (b optionsBuilder) Timeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "Timeout", value).(optionsBuilder)
}

func (b optionsBuilder) ConnectTimeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "ConnectTimeout", value).(optionsBuilder)
}

func (b optionsBuilder) Keyspace(value string) optionsBuilder {
	return builder.Set(b, "Keyspace", value).(optionsBuilder)
}

func (b optionsBuilder) NumConns(value int) optionsBuilder {
	return builder.Set(b, "NumConns", value).(optionsBuilder)
}

func (b optionsBuilder) Consistency(value string) optionsBuilder {
	return builder.Set(b, "Consistency", value).(optionsBuilder)
}

func (b optionsBuilder) SocketKeepalive(value time.Duration) optionsBuilder {
	return builder.Set(b, "SocketKeepalive", value).(optionsBuilder)
}

func (b optionsBuilder) MaxPreparedStmts(value int) optionsBuilder {
	return builder.Set(b, "MaxPreparedStmts", value).(optionsBuilder)
}

func (b optionsBuilder) MaxRoutingKeyInfo(value int) optionsBuilder {
	return builder.Set(b, "MaxRoutingKeyInfo", value).(optionsBuilder)
}

func (b optionsBuilder) PageSize(value int) optionsBuilder {
	return builder.Set(b, "PageSize", value).(optionsBuilder)
}

func (b optionsBuilder) DefaultTimestamp(value bool) optionsBuilder {
	return builder.Set(b, "DefaultTimestamp", value).(optionsBuilder)
}

func (b optionsBuilder) ReconnectInterval(value time.Duration) optionsBuilder {
	return builder.Set(b, "ReconnectInterval", value).(optionsBuilder)
}

func (b optionsBuilder) MaxWaitSchemaAgreement(value time.Duration) optionsBuilder {
	return builder.Set(b, "MaxWaitSchemaAgreement", value).(optionsBuilder)
}

func (b optionsBuilder) DisableInitialHostLookup(value bool) optionsBuilder {
	return builder.Set(b, "DisableInitialHostLookup", value).(optionsBuilder)
}

func (b optionsBuilder) WriteCoalesceWaitTime(value time.Duration) optionsBuilder {
	return builder.Set(b, "WriteCoalesceWaitTime", value).(optionsBuilder)
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
