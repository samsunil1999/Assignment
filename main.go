package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("")

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
