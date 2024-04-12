package main

import (
	"example.com/events-api/db"
	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":5050")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(400, gin.H{"message": "could not to fetch the data", "error": err})
		return
	}
	ctx.JSON(200, events)
}

func createEvents(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "could not parse the request data", "error": err})
		return
	}

	err = event.Save()

	if err != nil {
		ctx.JSON(500, gin.H{"message": "could not insert the data into database", "error": err})
		return
	}

	ctx.JSON(201, gin.H{"message": "Event Created successfully!", "data": event})
}
