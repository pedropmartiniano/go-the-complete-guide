package middlewares

import (
	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Not authorized"})
		return
	}

	c.Set("userId", userId)

	c.Next()
}
