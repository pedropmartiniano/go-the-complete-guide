package routes

import (
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the event Id", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not get the event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		c.JSON(500, gin.H{"message": "could not register user for event"})
		return
	}

	c.JSON(201, gin.H{"message": "Registered"})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the event Id", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not get the event"})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not cancel the registration"})
		return
	}

	c.JSON(200, gin.H{"message": "Canceled!"})
}

func getEventsRegistrations(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the event Id", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not find the event"})
		return
	}

	registrations, err := event.GetEventsRegistrations()

	if err != nil {
		c.JSON(400, gin.H{"message": "could not fetch the event registrations"})
		return
	}

	c.JSON(200, gin.H{"registrations": registrations})
}
