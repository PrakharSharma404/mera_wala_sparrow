package routes

import (
	"container-service/types"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// ----------- Prometheus metrics -----------

// Counter to track total build requests received
var buildRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "build_requests_total",
		Help: "Total number of Docker image build requests received",
	},
)

// Histogram to measure duration of build request handling
var buildRequestDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "build_request_duration_seconds",
		Help:    "Duration of Docker image build request handling in seconds",
		Buckets: prometheus.DefBuckets,
	},
)

func init() {
	// Register the metrics with Prometheus's default registry
	prometheus.MustRegister(buildRequestCounter)
	prometheus.MustRegister(buildRequestDuration)
}

// HandleBuildRequest builds docker image using given data
func HandleBuildRequest(c *gin.Context) {
	// Increment counter for each build request received
	buildRequestCounter.Inc()

	// Start timer for request duration measurement
	timer := prometheus.NewTimer(buildRequestDuration)
	defer timer.ObserveDuration()

	var request types.BuildRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	// Simulation
	time.Sleep(200 * time.Millisecond)
	imageTag := fmt.Sprintf("%s/%s", request.RepoOwner, request.RepoName)

	c.JSON(http.StatusOK, gin.H{
		"message":   "dummy build success",
		"image_tag": imageTag,
		"logs":      "dummy logs...",
	})
}
