package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var AmrPulls = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "arm_selector_arm_pulls",
		Help: "Total arm pulls.",
	},
	[]string{"experiment", "arm"},
)

func init() {
	prometheus.MustRegister(AmrPulls)
}
