package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var AmrPulls = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "arm_pulls",
		Help: "Total arm pulls.",
	},
	[]string{"experiment", "arm"},
)

func init() {
	// Register the custom counter metric with Prometheus
	prometheus.MustRegister(AmrPulls)
}
