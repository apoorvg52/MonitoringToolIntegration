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

func EventsHandler(c *gin.Context){

	// Path to the JSON file
    filePath := "C:/Users/apoor/OneDrive/Documents/ProjectJson/NewRelic/Event/event.json"

	compressedData, err := constants.ReadAndCompressJSON(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Metrics NOT Succesfully uploaded"})
    }

	res, err := services.SendEventsToNewRelic(constants.NEW_RELIC_API_KEY, compressedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Events NOT Succesfully uploaded", "Error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Metrics Succesfully uploaded", "uuid": res.Uuid})
}

func LogsHandler(c *gin.Context){

	var payload models.LogData
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send logs to new relic
	RequestId, err := services.SendLogsToNewRelic(constants.NEW_RELIC_API_KEY, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Logs NOT Succesfully uploaded", "Request Id": RequestId})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logs Succesfully uploaded", "Request Id": RequestId})
}

func TraceHandler(c *gin.Context){
	var traceReq models.TraceData
	// Convert trace data to JSON
    if err := c.ShouldBindJSON(&traceReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Upload Trace to NewRelic
	RequestId, err := services.SendTraceToNewRelic(constants.NEW_RELIC_API_KEY, traceReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Trace NOT Succesfully uploaded", "Request Id": RequestId})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trace Succesfully uploaded", "Request Id": RequestId})
}
