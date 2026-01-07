package handlers

import (
	"ManageTask/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//func ImportData(c *gin.Context) {
//	var task models.Task
//	c.ShouldBindJSON(&task)
//	fmt.Println(task.Name)
//}
//
//func Test(c *gin.Context) {
//	c.JSON(200, gin.H{"message": "Ngon"})
//}

func CheckOutTask(c *gin.Context) {
	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid task id"})
		return
	}

	type Ram struct {
		UserId int    `json:"user_id"`
		Value  string `json:"value"`
	}
	newRam := Ram{}
	c.ShouldBindJSON(&newRam)
	if newRam.UserId == 0 || newRam.Value == "" {
		c.JSON(400, gin.H{"error": "Info cannot be empty"})
		return
	}
	taskService := service.NewTaskService()

	task, err := taskService.GetTaskByUserAndId(newRam.UserId, int(id64))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	task.Status = newRam.Value
	taskService.UpdateTask(&task)
	c.JSON(200, gin.H{"message": "Task updated successfully"})

}

func ListTaskByUser(c *gin.Context) {
	type UserRam struct {
		UserId int `json:"user_id"`
	}

	userRam := UserRam{}

	c.ShouldBindJSON(&userRam)

	taskService := service.NewTaskService()
	listTask, err := taskService.GetTaskByUser(userRam.UserId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"tasks": listTask})
}
