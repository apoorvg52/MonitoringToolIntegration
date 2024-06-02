package handlers

import (
	// "monitoringtool/api/models"
	"monitoringtool/api/services"
	"monitoringtool/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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