package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ReturnJsonStruct(c *gin.Context, genericStruct interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(c.Writer).Encode(genericStruct)
}
