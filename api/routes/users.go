package routes

import (
	"example.com/events-api/models"
	"example.com/events-api/utils"
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

func signin(c *gin.Context) {
	var signinBody models.User
	err := c.ShouldBindJSON(&signinBody)

	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the request body data"})
		return
	}

	user, err := models.GetUserByEmail(signinBody.Email)

	if err != nil {
		c.JSON(401, gin.H{"message": "email or password incorrect"})
		return
	}

	isPasswordCorrect := utils.CompareHashPassword(user.Password, signinBody.Password)

	if !isPasswordCorrect {
		c.JSON(401, gin.H{"message": "email or password incorrect"})
		return
	}

	token, err := utils.GenerateJwtToken(user.Email, user.Id)

	if err != nil {
		c.JSON(401, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(200, gin.H{"message": "user authenticated", "token": token})
}
