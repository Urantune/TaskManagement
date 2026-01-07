package main

import (
	"ManageTask/handlers"
	"ManageTask/middleware"
	"ManageTask/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := repository.Connect(); err != nil {
		log.Fatal("DB connect failed:", err)
	}

	r := gin.Default()

	r.POST("/login", handlers.Login)

	r.POST("/register", handlers.Register)

	auth := r.Group("/")
	auth.Use(middleware.MidwareAuth())

	//	auth.POST("/test", handlers.Test)

	admin := r.Group("/admin")

	admin.Use(
		middleware.MidwareAuth(),
		middleware.CheckRole("admin"),
	)

	//	admin.POST("/adminTest", handlers.TestAdminTask)

	admin.POST("/addTask", handlers.CreateTask)

	admin.POST("/listTask", handlers.ListTask)

	admin.POST("/editTask/:id", handlers.EditTask)

	admin.POST("/deleteTask/:id", handlers.DeleteTask)

	r.Run(":8080")

}
