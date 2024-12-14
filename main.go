package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var cache = make(map[uuid.UUID]int64)

func main() {
	router := gin.Default()
	router.GET("/receipts/{id}/points", getPoints)
	router.POST("/receipts/process", postReceipt)
	router.Run("localhost:8080")
}

// receipts/{id}/points
func getPoints(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	points, ok := cache[uuid]

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, "ID "+id+" not in cache")
	}
	c.IndentedJSON(http.StatusOK, points)
}

// /receipts/process
func postReceipt(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"id": ""})
}
