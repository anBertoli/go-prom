package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var endpointCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "endpoint_count",
		Help: "endpoint count",
	},
	[]string{"path", "code"},
)

var latencyHist = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "latency_microseconds",
		Help:    "latency of requests in microseconds",
		Buckets: []float64{0.1, 1, 10, 100, 250, 500, 1000, 2500, 5000, 10000},
	},
	[]string{"path"},
)

var tempGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "fake_temp",
		Help: "temperature (fake)",
	},
	nil,
)

func init() {
	must := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	must(prometheus.Register(endpointCount))
	must(prometheus.Register(latencyHist))
	must(prometheus.Register(tempGauge))
}

func recordTemperature() {
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		tempGauge.WithLabelValues().Set(rand.Float64() * 20)
	}
}
