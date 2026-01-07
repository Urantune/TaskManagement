package handlers

import (
	"ManageTask/models"
	"ManageTask/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//func TestAdminTask(c *gin.Context) {
//
//	c.JSON(200, gin.H{"message": "NgonAdmin"})
//}

func CreateTask(c *gin.Context) {
	task := models.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if task.Name == "" || task.Status == "" {
		c.JSON(400, gin.H{"error": "Info cannot be empty"})
		return
	}
	taskService := service.NewTaskService()
	if err := taskService.AddTask(&task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Task created successfully"})

}

func ListTask(c *gin.Context) {
	taskService := service.NewTaskService()
	listTask, err := taskService.GetTasks()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tasks": listTask})
}

func EditTask(c *gin.Context) {

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid task id"})
		return
	}

	type Ram struct {
		Name string `json:"name"`
	}
	newRam := Ram{}

	c.ShouldBindJSON(&newRam)
	taskService := service.NewTaskService()
	task, err := taskService.GetOneTask(int(id64))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	task.Name = newRam.Name
	if newRam.Name == "" {
		c.JSON(400, gin.H{"error": "Info cannot be empty"})
		return
	}

	err = taskService.UpdateTask(&task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task updated successfully"})

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid task id"})
		return
	}
	taskService := service.NewTaskService()
	err = taskService.DeleteTask(int(id64))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
