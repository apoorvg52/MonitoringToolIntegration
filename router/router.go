package router

import (
	"github.com/gin-gonic/gin"
	"monitoringtool/api/handlers"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/uploadMetric", handlers.MetricsHandler)
	r.POST("/uploadEvents", handlers.EventsHandler)
	r.POST("/uploadLogs", handlers.LogsHandler)
	r.POST("/uploadTrace", handlers.TraceHandler)
}
