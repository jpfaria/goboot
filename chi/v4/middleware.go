package gichi

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/b2wdigital/goignite/info"
	gilog "github.com/b2wdigital/goignite/log"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	nr "github.com/newrelic/go-agent/v3/newrelic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	uuid "github.com/satori/go.uuid"
)

// NewTidMiddleware is a middleware that looks for a XTID value inside the http.Request
// and generate one if it does not exists.
func NewTidMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tid := r.Header.Get("x-tid")
			if tid == "" {
				tid = info.AppName + "-" + uuid.NewV4().String()
			}
			w.Header().Set("x-tid", tid)
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// NewLogMiddleware returns a middleware that logs HTTP requests.
func NewLogMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		t1 := time.Now()
		reqId := middleware.GetReqID(ctx)
		preReqContent := gilog.Fields{
			"time":      t1,
			"requestId": reqId,
			"method":    r.Method,
			"endpoint":  r.RequestURI,
			"protocol":  r.Proto,
		}

		if r.RemoteAddr != "" {
			preReqContent["ip"] = r.RemoteAddr
		}

		tid := r.Header.Get("X-TID")
		if tid != "" {
			preReqContent["tid"] = tid
		}

		logger := gilog.FromContext(ctx).WithFields(preReqContent)
		ctx = logger.ToContext(ctx)
		r = r.WithContext(ctx)
		logger.Info("request started")

		defer func() {
			if err := recover(); err != nil {
				gilog.WithFields(
					gilog.Fields{
						"requestId":  reqId,
						"duration":   time.Since(t1),
						"status":     500,
						"stacktrace": string(debug.Stack()),
					},
				).Error("request finished with panic")
				panic(err)
			}
		}()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		status := ww.Status()
		postReqContent := gilog.Fields{
			"requestId":     reqId,
			"duration":      time.Since(t1),
			"contentLength": ww.BytesWritten(),
			"status":        status,
		}

		if cache := ww.Header().Get("x-cache"); cache != "" {
			postReqContent["cache"] = cache
		}

		logger = gilog.WithFields(postReqContent)
		if status >= 100 && status < 400 {
			logger.Info("request finished")
		} else if status == 500 {
			logger.WithField("stacktrace",
				string(debug.Stack())).Info("internal error during request")
		} else {
			message := "request finished"

			// FIX: For some reason, the 'context.deadlineExceededError{}' isn't getting into here, we
			// did a quick fix checking the status code and returing the same message as the error., but
			// something is wrong and we need fix it.
			if status == 504 {
				message += ": context deadline exceeded"
			} else {
				if err := ctx.Err(); err != nil {
					message += fmt.Sprintf(": %s", err.Error())
				}
			}
			logger.Error(message)
		}
	}

	return http.HandlerFunc(fn)
}

func NewNewRelicMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		url := r.URL.String()
		path := r.URL.Path
		txnName := strings.Join([]string{r.Method, path}, " ")

		txn := ginewrelic.Application().StartTransaction(txnName)
		defer txn.End()

		txn.SetWebRequestHTTP(r)

		if GetNewRelicWebResponseEnabled() {
			w = txn.SetWebResponse(w)
		}

		txn.AddAttribute("request.url", fmt.Sprintf("http://%s%s", r.Host, url))

		qs := r.URL.Query()
		for key, value := range qs {
			txn.AddAttribute(key, strings.Join(value, "|"))
		}

		if reqID := middleware.GetReqID(ctx); reqID != "" {
			txn.AddAttribute("request.id", reqID)
		}

		r = nr.RequestWithTransactionContext(r, txn)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// NewMetricMiddleware is a prometheus
func NewMetricMiddleware(next http.Handler) http.Handler {

	const (
		httpRequestsCount    = "requests_total"
		httpRequestsDuration = "request_duration_seconds"
	)

	config := struct {
		Namespace string
		Buckets   []float64
		Subsystem string
	}{
		Namespace: "chi",
		Subsystem: "http",
		Buckets: []float64{
			0.0005,
			0.001, // 1ms
			0.002,
			0.005,
			0.01, // 10ms
			0.02,
			0.05,
			0.1, // 100 ms
			0.2,
			0.5,
			1.0, // 1s
			2.0,
			5.0,
			10.0, // 10s
			15.0,
			20.0,
			30.0,
		},
	}

	normalizeHTTPStatus := func(status int) string {
		if status < 200 {
			return "1xx"
		} else if status < 300 {
			return "2xx"
		} else if status < 400 {
			return "3xx"
		} else if status < 500 {
			return "4xx"
		}

		return "5xx"
	}

	httpRequests := promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: config.Namespace,
		Subsystem: config.Subsystem,
		Name:      httpRequestsCount,
		Help:      "Number of HTTP operations",
	}, []string{"status", "method", "path"})

	httpDuration := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: config.Namespace,
		Subsystem: config.Subsystem,
		Name:      httpRequestsDuration,
		Help:      "Spend time by processing a route",
		Buckets:   config.Buckets,
	}, []string{"status", "method", "path"})

	fn := func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		route := chi.RouteContext(ctx)
		path := route.RoutePatterns[0]
		status := normalizeHTTPStatus(ww.Status())

		httpRequests.WithLabelValues(status, r.Method, path).Inc()
		prometheus.NewTimer(httpDuration.WithLabelValues(status, r.Method, path)).ObserveDuration()

	}

	return http.HandlerFunc(fn)
}
