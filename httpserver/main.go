package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/about", AboutHandler)
	mux.HandleFunc("/contact", ContactHandler)
	mux.Handle("/metrics", promhttp.Handler())

	go recordTemperature()

	err := http.ListenAndServe("127.0.0.1:4000", mux)
	log.Println("server stopped", err)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	httpCode := randHTTPCode()

	defer func(t time.Time) {
		endpointCount.WithLabelValues(r.URL.Path, fmt.Sprintf("%d", httpCode)).Inc()
		latencyHist.WithLabelValues(r.URL.Path).Observe(float64(time.Since(t).Microseconds()))
	}(time.Now())

	// simulate some latency
	time.Sleep(1 * time.Microsecond)
	w.WriteHeader(httpCode)
	w.Write([]byte("<h1>Home Page</h1>\n"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	httpCode := randHTTPCode()

	defer func(t time.Time) {
		endpointCount.WithLabelValues(r.URL.Path, fmt.Sprintf("%d", httpCode)).Inc()
		latencyHist.WithLabelValues(r.URL.Path).Observe(float64(time.Since(t).Microseconds()))
	}(time.Now())

	// simulate some latency
	time.Sleep(2 * time.Microsecond)
	w.WriteHeader(httpCode)
	w.Write([]byte("<h1>About Page</h1>\n"))
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	httpCode := randHTTPCode()

	defer func(t time.Time) {
		endpointCount.WithLabelValues(r.URL.Path, fmt.Sprintf("%d", httpCode)).Inc()
		latencyHist.WithLabelValues(r.URL.Path).Observe(float64(time.Since(t).Microseconds()))
	}(time.Now())

	// simulate some latency
	time.Sleep(5 * time.Microsecond)
	w.WriteHeader(httpCode)
	w.Write([]byte("<h1>Contact Page</h1>\n"))
}

func randHTTPCode() int {
	switch rand.Intn(5) {
	case 0:
		return http.StatusBadRequest
	case 1:
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}
