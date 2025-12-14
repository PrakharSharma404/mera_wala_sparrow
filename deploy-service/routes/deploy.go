package routes

import (
	"deploy-service/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// deploy given docker image
func HandleDeploy(c *gin.Context) {
	// -----------------------------------------------------------------
	// What: Start a timer to measure deployment duration.
	// Why:  To capture the exact time taken to process the deployment request.
	// How:  prometheus.NewTimer creates a timer that observes the duration
	//       when ObserveDuration() is called (via defer).
	// -----------------------------------------------------------------
	timer := prometheus.NewTimer(deployDuration)
	defer timer.ObserveDuration()

	// -----------------------------------------------------------------
	// What: Increment the deployment counter with status "started".
	// Why:  To track the number of initiated deployments.
	// How:  Access the counter vector with label "started" and call Inc().
	// -----------------------------------------------------------------
	deployCounter.WithLabelValues("started").Inc()

	var request types.DeployRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// ---------------------------------------------------------
		// What: Record a deployment failure.
		// Why:  To track failed attempts due to bad requests.
		// How:  Increment counter with label "failure".
		// ---------------------------------------------------------
		deployCounter.WithLabelValues("failure").Inc()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	// Logic for deployment would go here...

	// ---------------------------------------------------------
	// What: Record a successful deployment.
	// Why:  To complete the logic and track success rate.
	// How:  Increment counter with label "success".
	// ---------------------------------------------------------
	deployCounter.WithLabelValues("success").Inc()
	c.JSON(http.StatusOK, gin.H{"message": "deployment started"})
}
