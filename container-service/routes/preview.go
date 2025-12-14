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

// Counter for total preview requests received
var previewRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "preview_requests_total",
		Help: "Total number of Dockerfile preview requests received",
	},
)

// Histogram for measuring duration of preview request processing
var previewRequestDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "preview_request_duration_seconds",
		Help:    "Duration of Dockerfile preview request processing in seconds",
		Buckets: prometheus.DefBuckets,
	},
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(previewRequestCounter)
	prometheus.MustRegister(previewRequestDuration)
}

// HandlePreviewRequest generates a Dockerfile preview based on query params
func HandlePreviewRequest(c *gin.Context) {
	previewRequestCounter.Inc()             // Increment request count
	timer := prometheus.NewTimer(previewRequestDuration) // Start timer
	defer timer.ObserveDuration()            // Observe duration at return

	var params types.PreviewRequest
	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	// Simulation
	time.Sleep(100 * time.Millisecond)

	dockerfile := fmt.Sprintf("FROM dummy-image\nRUN echo 'This is a dummy preview for %s'", params.ProjectType)

	c.String(http.StatusOK, dockerfile)
}
