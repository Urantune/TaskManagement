package service

import (
	"ManageTask/models"
	"ManageTask/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepo
}

func NewTaskService() *TaskService {
	return &TaskService{
		taskRepo: repository.NewTaskRepository(repository.DB),
	}
}

func (r *TaskService) AddTask(task *models.Task) error {
	return r.taskRepo.AddTask(task)
}

func (r *TaskService) GetTasks() ([]models.Task, error) {
	return r.taskRepo.GetTasks()
}

func (r *TaskService) GetOneTask(id int) (models.Task, error) {
	return r.taskRepo.GetOneTask(id)
}

func (r *TaskService) UpdateTask(task *models.Task) error {
	return r.taskRepo.UpdateTask(task)
}

func (r *TaskService) DeleteTask(id int) error {
	return r.taskRepo.DeleteTask(id)
}
