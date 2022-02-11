package main

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const AppName = "test_app"

const (
	labelApp    = "app"
	labelPath   = "path"
	labelCode   = "code"
	labelMethod = "method"
)

var (
	httpDurationSummary   *prometheus.SummaryVec
	httpDurationHistogram *prometheus.HistogramVec
	httpCount             *prometheus.CounterVec
	httpGauge             *prometheus.GaugeVec
)

func InitMetrics() {
	if httpDurationSummary == nil {
		httpDurationSummary = promauto.NewSummaryVec(prometheus.SummaryOpts{
			Name:       "http_request_duration_summary",
			Help:       "Duration of HTTP requests Summary.",
			Objectives: map[float64]float64{0.5: 0.5, 0.9: 0.9, 1: 1},
			AgeBuckets: 3,
			MaxAge:     120 * time.Second,
		}, []string{labelApp, labelPath, labelCode, labelMethod})
	}

	if httpDurationHistogram == nil {
		httpDurationHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_request_duration_histogram",
			Help:    "Duration of HTTP requests Histogram.",
			Buckets: []float64{10, 50, 100, 200, 500, 1000},
		}, []string{labelApp, labelPath, labelCode, labelMethod})
	}

	if httpCount == nil {
		httpCount = promauto.NewCounterVec(prometheus.CounterOpts{
			Name:        "http_request_count",
			Help:        "Count of HTTP requests.",
			ConstLabels: nil,
		}, []string{labelApp, labelPath, labelCode, labelMethod})
	}

	if httpGauge == nil {
		httpGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "http_request_gauge",
			Help:        "Gauge of HTTP requests.",
			ConstLabels: nil,
		}, []string{labelApp, labelPath, labelCode, labelMethod})
	}
}

func SaveHTTPDuration(timeSince time.Time, path string, code int, method string) {
	httpDurationSummary.With(map[string]string{
		labelApp:    AppName,
		labelPath:   path,
		labelCode:   strconv.Itoa(code),
		labelMethod: method,
	}).Observe(float64(time.Since(timeSince).Milliseconds()))
}

func SaveHTTPDurationHistogram(timeSince time.Time, path string, code int, method string) {
	httpDurationHistogram.With(map[string]string{
		labelApp:    AppName,
		labelPath:   path,
		labelCode:   strconv.Itoa(code),
		labelMethod: method,
	}).Observe(float64(time.Since(timeSince).Milliseconds()))
}

func SaveHTTPCount(value float64, path string, code int, method string) {
	httpCount.With(map[string]string{
		labelApp:    AppName,
		labelPath:   path,
		labelCode:   strconv.Itoa(code),
		labelMethod: method,
	}).Add(value)
}

func SaveHTTPGauge(value float64, path string, code int, method string) {
	httpGauge.With(map[string]string{
		labelApp:    AppName,
		labelPath:   path,
		labelCode:   strconv.Itoa(code),
		labelMethod: method,
	}).Set(value)
}