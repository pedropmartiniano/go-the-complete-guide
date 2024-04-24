package routes

import (
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(400, gin.H{"message": "could not to fetch the data", "error": err})
		return
	}
	c.JSON(200, events)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the event Id", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not get the event by his id", "error": err})
		return
	}

	c.JSON(200, gin.H{"data": event})
}

func createEvents(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the request data", "error": err})
		return
	}

	userId := c.GetInt64("userId")
	event.UserId = userId
	err = event.Save()

	if err != nil {
		c.JSON(500, gin.H{"message": "could not insert the data into database", "error": err})
		return
	}

	c.JSON(201, gin.H{"message": "Event Created successfully!", "data": event})
}

func updateEvent(c *gin.Context) {
	var eventBody models.Event
	c.ShouldBindJSON(&eventBody)

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	userId := c.GetInt64("userId")

	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the event Id", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not get the event", "error": err})
		return
	}

	if event.UserId != userId {
		c.JSON(403, gin.H{"message": "you can only update your our events"})
		return
	}

	eventBody.Id = eventId

	err = eventBody.Update()

	if err != nil {
		c.JSON(400, gin.H{"message": "could not update event", "error": err})
		return
	}

	c.JSON(200, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	userId := c.GetInt64("userId")

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not get the event"})
		return
	}

	if userId != event.UserId {
		c.JSON(403, gin.H{"message": "you can only delete your our events"})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not delete the event"})
		return
	}

	c.JSON(200, gin.H{"message": "Event deleted successfully"})
}
