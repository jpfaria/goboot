package datadog

import (
	"net"
	"os"

	"github.com/b2wdigital/goignite/v2/contrib/net/http/client"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/info"
)

type Options struct {
	Service       string
	Env           string
	Enabled       bool
	Tags          map[string]string
	Host          string
	Port          string
	LambdaMode    bool
	Analytics     bool
	AnalyticsRate float64
	DebugMode     bool
	DebugStack    bool
	HttpClient    client.Options
	Version       string
	Log           struct {
		Level string
	}
	Addr string
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

func NewOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	if v := info.AppName; v != "" {
		o.Service = v
	}

	if v := os.Getenv("DD_SERVICE"); v != "" {
		o.Service = v
	}

	if v := os.Getenv("DD_AGENT_HOST"); v != "" {
		o.Host = v
	}

	if v := os.Getenv("DD_TRACE_AGENT_PORT"); v != "" {
		o.Port = v
	}

	if v := os.Getenv("DD_ENV"); v != "" {
		o.Env = v
	}

	if v := info.Version; v != "" {
		o.Version = v
	}

	if v := os.Getenv("DD_VERSION"); v != "" {
		o.Version = v
	}

	o.Addr = net.JoinHostPort(o.Host, o.Port)

	return o, nil
}
