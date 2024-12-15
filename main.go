package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var cache = make(map[uuid.UUID]int)

func main() {
	router := gin.Default()
	router.GET("/receipts/:id/points", getPoints)
	router.POST("/receipts/process", postReceipt)
	router.Run("localhost:8080")
}

// GET receipts/:id/points
func getPoints(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "ID not found in cache")
		return
	}

	points, ok := cache[uuid]

	if !ok {
		c.IndentedJSON(http.StatusNotFound, "ID not found in cache")
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"points": points})
}

// POST /receipts/process
func postReceipt(c *gin.Context) {
	receipt := Receipt{}

	if err := c.BindJSON(&receipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Receipt is invalid")
		return
	}

	valid := validateReceipt(receipt)

	if !valid {
		c.IndentedJSON(http.StatusBadRequest, "Receipt is invalid")
		return
	}

	points := calculatePoints(receipt)

	uuid := uuid.New()
	cache[uuid] = points

	c.IndentedJSON(http.StatusOK, gin.H{"id": uuid})
}
