package main

import (
	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":5050")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(200, events)
}

func createEvents(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(400, gin.H{"message": "could not parse the request data", "error": err})
		return
	}

	event.Id = 1
	event.UserId = 1

	ctx.JSON(201, gin.H{"message": "Event Created successfully!", "data": event})
}
