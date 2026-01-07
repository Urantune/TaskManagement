package handlers

import "github.com/gin-gonic/gin"

func TestAdminTask(c *gin.Context) {

	c.JSON(200, gin.H{"message": "NgonAdmin"})
}
