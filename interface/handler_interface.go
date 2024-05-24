package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	UploadCSVHandler(c *gin.Context)
}