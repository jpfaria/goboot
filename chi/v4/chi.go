package gichi

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	instance *chi.Mux
)

func NewMux(ctx context.Context) *chi.Mux {

	instance = chi.NewRouter()

	setDefaultMiddlewares(ctx, instance)
	setDefaultRouters(ctx, instance)

	return instance
}

func setDefaultMiddlewares(ctx context.Context, instance *chi.Mux) {

	if GetMiddlewareRecoverEnabled() {
		instance.Use(middleware.Recoverer)
	}
	if GetMiddlewareRealIPEnabled() {
		instance.Use(middleware.RealIP)
	}
	if GetMiddlewareRequestIDEnabled() {
		instance.Use(middleware.RequestID)
	}
	if GetMiddlewareNewTidEnabled() {
		instance.Use(NewTidMiddleware())
	}
	if GetMiddlewareNewRelicEnabled() {
		instance.Use(NewNewRelicMiddleware)
	}
	if GetMiddlewareLoggerEnabled() {
		instance.Use(NewLogMiddleware)
	}
	if GetMiddlewareMetricEnabled() {
		instance.Use(NewMetricMiddleware)
	}
}

func setDefaultRouters(ctx context.Context, instance *chi.Mux) {

	l := gilog.FromContext(ctx)

	statusRoute := GetStatusRoute()

	l.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.Get(GetStatusRoute(), statusHandler.Get())

	healthRoute := GetHealthRoute()

	l.Infof("configuring health router on %s", healthRoute)

	healthHandler := NewHealthHandler()

	instance.Get(healthRoute, healthHandler.Get(ctx))

	if GetMiddlewareMetricEnabled() {

		metricRoute := GetMetricRoute()

		l.Infof("configuring metric router on %s", metricRoute)

		instance.Handle(metricRoute, promhttp.Handler())
	}

}
