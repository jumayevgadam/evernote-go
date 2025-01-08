package metrics

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// App Metrics interface.
type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
}

// Prometheus metrics struct.
type PrometheusMetrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

// CreateMetrics with address and name.
func CreateMetrics(address, name string) (Metrics, error) {
	var metr PrometheusMetrics
	metr.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_hits_total",
	})

	if err := prometheus.Register(metr.HitsTotal); err != nil {
		return nil, fmt.Errorf("prometheus.Register: metr.HitsTotal=> %w", err)
	}

	metr.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name + "_hits",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Hits); err != nil {
		return nil, fmt.Errorf("prometheus.Register: metr.Hits => %w", err)
	}

	metr.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name + "_times",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Times); err != nil {
		return nil, fmt.Errorf("prometheus.Register: metr.Times => %w", err)
	}

	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		return nil, fmt.Errorf("prometheus.Register: prometheus.NewBuildInfoCollector =>%w", err)
	}

	go func() {
		r := gin.New()
		r.GET("/metrics", gin.WrapH(promhttp.Handler()))

		if err := r.Run(address); err != nil {
			zap.L().Error("error in running metrics")
		}
	}()

	return &metr, nil
}

func (metr *PrometheusMetrics) IncHits(status int, method, path string) {
	metr.HitsTotal.Inc()
	metr.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

func (metr *PrometheusMetrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	metr.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}
