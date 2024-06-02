package handlers

import (
	"monitoringtool/api/models"
	"monitoringtool/api/services"
	"monitoringtool/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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