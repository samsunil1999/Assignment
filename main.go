package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {

	// mapping of endpoints
	mapUrls()

	// run the application
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}

}
