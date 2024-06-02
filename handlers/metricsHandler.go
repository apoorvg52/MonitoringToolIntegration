package handlers

import (
	"monitoringtool/api/models"
	"monitoringtool/api/services"
	"monitoringtool/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MetricsHandler(c *gin.Context) {
	var payload []models.Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process Payload & unix time stamp
	metricsPayload := constants.ProcessPayload(payload)

	// Upload Metrics to NewRelic
	RequestId, err := services.SendMetricToNewRelic(constants.NEW_RELIC_API_KEY, metricsPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Metrics NOT Succesfully uploaded", "Request Id": RequestId})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Metrics Succesfully uploaded", "Request Id": RequestId})
}