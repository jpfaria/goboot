package datadog

import (
	"context"
	"sync"

	"github.com/b2wdigital/goignite/v2/contrib/net/http/client"
	"github.com/b2wdigital/goignite/v2/core/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var once sync.Once

func NewTracer(ctx context.Context, options *Options, startOptions ...tracer.StartOption) {

	if !IsEnabled() {
		return
	}

	once.Do(func() {

		logger := log.FromContext(ctx)

		httpClient := client.NewClient(ctx, &options.HttpClient)

		so := []tracer.StartOption{
			tracer.WithAgentAddr(options.Addr),
			tracer.WithEnv(options.Env),
			tracer.WithService(options.Service),
			tracer.WithServiceVersion(options.Version),
			tracer.WithLogger(NewLogger()),
			tracer.WithHTTPClient(httpClient),
			tracer.WithAnalytics(options.Analytics),
			tracer.WithAnalyticsRate(options.AnalyticsRate),
			tracer.WithLambdaMode(options.LambdaMode),
			tracer.WithDebugMode(options.DebugMode),
			tracer.WithDebugStack(options.DebugStack),
		}

		for k, v := range options.Tags {
			so = append(so, tracer.WithGlobalTag(k, v))
		}

		so = append(so, startOptions...)

		tracer.Start(so...)

		logger.Infof("started a datadog tracer: %s", options.Service)
	})

}
