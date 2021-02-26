package middleware

// from github.com/labstack/echo-contrib/prometheus
// with some updates:
// - fix histogram_vec to include Buckets
// - remove subsystem https://github.com/prometheus/client_golang/issues/240

import (
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var defaultMetricPath = "/metrics"

// Standard default metrics
//	counter, counter_vec, gauge, gauge_vec,
//	histogram, histogram_vec, summary, summary_vec
// DefBuckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
var reqCnt = &Metric{
	ID:          "reqCnt",
	Name:        "echo_requests_total",
	Description: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	Args:        []string{"service", "code", "method", "url", "usertype"},
	Type:        "counter_vec",
}

var reqDur = &Metric{
	ID:          "reqDur",
	Name:        "echo_request_duration_seconds",
	Description: "The HTTP request latencies in seconds.",
	Args:        []string{"service", "code", "method", "url", "usertype"},
	Type:        "histogram_vec",
	Buckets:     prometheus.DefBuckets,
}

var resSz = &Metric{
	ID:          "resSz",
	Name:        "echo_response_size_bytes",
	Description: "The HTTP response sizes in bytes.",
	Args:        []string{"service", "code", "method", "url", "usertype"},
	Type:        "histogram_vec",
	Buckets:     prometheus.LinearBuckets(0, 1000, 10),
}

var reqSz = &Metric{
	ID:          "reqSz",
	Name:        "echo_request_size_bytes",
	Description: "The HTTP request sizes in bytes.",
	Args:        []string{"service", "code", "method", "url", "usertype"},
	Type:        "histogram_vec",
	Buckets:     prometheus.LinearBuckets(0, 1000, 10),
}

var standardMetrics = []*Metric{
	reqCnt,
	reqDur,
}

/*
RequestCounterURLLabelMappingFunc is a function which can be supplied to the middleware to control
the cardinality of the request counter's "url" label, which might be required in some contexts.
For instance, if for a "/customer/:name" route you don't want to generate a time series for every
possible customer name, you could use this function:

func(c echo.Context) string {
	url := c.Request.URL.Path
	for _, p := range c.Params {
		if p.Key == "name" {
			url = strings.Replace(url, p.Value, ":name", 1)
			break
		}
	}
	return url
}

which would map "/customer/alice" and "/customer/bob" to their template "/customer/:name".
*/
type RequestCounterURLLabelMappingFunc func(c echo.Context) string

// Metric is a definition for the name, description, type, ID, and
// prometheus.Collector type (i.e. CounterVec, Summary, etc) of each metric
type Metric struct {
	MetricCollector prometheus.Collector
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
	Buckets         []float64
}

// Prometheus contains the metrics gathered by the instance and its path
type Prometheus struct {
	reqCnt               *prometheus.CounterVec
	reqDur, reqSz, resSz *prometheus.HistogramVec
	router               *echo.Echo
	listenAddress        string

	MetricsList []*Metric
	MetricsPath string
	Subsystem   string
	Skipper     middleware.Skipper

	RequestCounterURLLabelMappingFunc RequestCounterURLLabelMappingFunc
}

// NewPrometheus generates a new set of metrics with a certain subsystem name
func NewPrometheus(subsystem string, customMetricsList ...[]*Metric) *Prometheus {
	var metricsList []*Metric

	if len(customMetricsList) > 1 {
		panic("Too many args. NewPrometheus( string, <optional []*Metric> ).")
	} else if len(customMetricsList) == 1 {
		metricsList = customMetricsList[0]
	}

	metricsList = append(metricsList, standardMetrics...)

	p := &Prometheus{
		MetricsList: metricsList,
		MetricsPath: defaultMetricPath,
		Subsystem:   subsystem,
		Skipper:     DefaultSkipper,
		RequestCounterURLLabelMappingFunc: func(c echo.Context) string {
			url := c.Path()
			for _, p := range c.ParamNames() {
				if p == "id" {
					url = strings.Replace(url, c.Param(p), ":id", 1)
					break
				}
			}
			return url
			//return c.Path() // i.e. by default do nothing, i.e. return URL as is
		},
	}

	p.registerMetrics(subsystem)

	return p
}

// SetMetricsPath set metrics paths
func (p *Prometheus) SetMetricsPath(e *echo.Echo) {
	if p.listenAddress != "" {
		p.router.GET(p.MetricsPath, prometheusHandler())
		p.runServer()
	} else {
		e.GET(p.MetricsPath, prometheusHandler())
	}
}

func (p *Prometheus) routerStart(listenAddress string) {
	err := p.router.Start(listenAddress)
	if err != nil {
		panic("prometheus router start failed")
	}
}

func (p *Prometheus) runServer() {
	if p.listenAddress != "" {
		go p.routerStart(p.listenAddress)
	}
}

// NewMetric associates prometheus.Collector based on Metric.Type
func NewMetric(m *Metric, subsystem string) prometheus.Collector {
	var metric prometheus.Collector
	switch m.Type {
	case "counter_vec":
		metric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: m.Name,
				Help: m.Description,
			},
			m.Args,
		)
	case "counter":
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	case "gauge_vec":
		metric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: m.Name,
				Help: m.Description,
			},
			m.Args,
		)
	case "gauge":
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	case "histogram_vec":
		metric = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    m.Name,
				Help:    m.Description,
				Buckets: m.Buckets,
			},
			m.Args,
		)
	case "histogram":
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	case "summary_vec":
		metric = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: m.Name,
				Help: m.Description,
			},
			m.Args,
		)
	case "summary":
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	}
	return metric
}

func (p *Prometheus) registerMetrics(subsystem string) {

	for _, metricDef := range p.MetricsList {
		metric := NewMetric(metricDef, subsystem)
		if err := prometheus.Register(metric); err != nil {
			log.Errorf("%s could not be registered in Prometheus: %v", metricDef.Name, err)
		}
		switch metricDef {
		case reqCnt:
			p.reqCnt = metric.(*prometheus.CounterVec)
		case reqDur:
			p.reqDur = metric.(*prometheus.HistogramVec)
		case resSz:
			p.resSz = metric.(*prometheus.HistogramVec)
		case reqSz:
			p.reqSz = metric.(*prometheus.HistogramVec)
		}
		metricDef.MetricCollector = metric
	}
}

// Use adds the middleware to the Echo engine.
func (p *Prometheus) Use(e *echo.Echo) {
	e.Use(p.HandlerFunc)
	p.SetMetricsPath(e)
}

// HandlerFunc defines handler function for middleware
func (p *Prometheus) HandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if c.Path() == p.MetricsPath {
			return next(c)
		}
		if p.Skipper(c) {
			return next(c)
		}

		start := time.Now()

		if err = next(c); err != nil {
			c.Error(err)
		}

		stop := time.Now()

		userType, ok := c.Get("userType").(string)
		if !ok {
			userType = "unknown"
		}
		if len(userType) == 0 {
			userType = "unknown"
		}

		status := strconv.Itoa(c.Response().Status)
		url := p.RequestCounterURLLabelMappingFunc(c)

		elapsed := float64(stop.Sub(start).Nanoseconds() / int64(time.Second))

		p.reqDur.WithLabelValues(p.Subsystem, status, c.Request().Method, url, userType).Observe(elapsed)
		p.reqCnt.WithLabelValues(p.Subsystem, status, c.Request().Method, url, userType).Inc()

		return
	}
}

func prometheusHandler() echo.HandlerFunc {
	h := promhttp.Handler()
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
