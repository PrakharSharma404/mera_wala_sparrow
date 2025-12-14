package routes

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ---------------------------------------------------------------------
// What: Define a Counter Vector for tracking deployments.
// Why:  To filter deployments by status (success vs failure).
//       Critical for monitoring deployment reliability.
// How:  Creates a Prometheus CounterVec with a "status" label.
// ---------------------------------------------------------------------
var deployCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "deployments_total",
		Help: "Total number of deployments attempted",
	},
	[]string{"status"}, // "success", "failure"
)

// ---------------------------------------------------------------------
// What: Define a Histogram for tracking deployment duration.
// Why:  To measure how long deployments take, identifying performance bottlenecks.
// How:  Creates a Prometheus Histogram with default buckets.
// ---------------------------------------------------------------------
var deployDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "deployment_duration_seconds",
		Help:    "Duration of deployments in seconds",
		Buckets: prometheus.DefBuckets,
	},
)

// ---------------------------------------------------------------------
// What: Initialize deployment metrics and start the HTTP server.
// Why:  To make metrics available for scraping by Prometheus.
// How:  Registers metrics and listens on port 2112 for /metrics.
// ---------------------------------------------------------------------
func InitMetrics() {
	// Register metrics
	prometheus.MustRegister(deployCounter)
	prometheus.MustRegister(deployDuration)

	// Start metrics server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Starting Prometheus metrics server on :2112/metrics")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()
}
