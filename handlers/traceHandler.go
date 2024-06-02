package handlers

import (
	"monitoringtool/api/models"
	"monitoringtool/api/services"
	"monitoringtool/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
