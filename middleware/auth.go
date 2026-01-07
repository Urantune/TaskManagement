package middleware

import (
	"ManageTask/utils"

	"github.com/gin-gonic/gin"
)

func MidwareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		claim, err := utils.ParseToken(token)
		if err != nil {

			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		c.Set("user", claim)
		c.Next()
	}
}

func CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, ok := c.Get("user")

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		claim, ok := user.(*utils.Claims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if claim.UserRole != role {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
