package routes

import (
	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request body data"})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not create user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully", "data": user})
}
