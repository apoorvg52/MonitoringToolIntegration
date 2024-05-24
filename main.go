package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"monitoringtool/api/router"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
