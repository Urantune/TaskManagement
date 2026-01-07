package handlers

import (
	"ManageTask/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ImportData(c *gin.Context) {
	var task models.Task
	c.ShouldBindJSON(&task)
	fmt.Println(task.Name)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Ngon"})
}
