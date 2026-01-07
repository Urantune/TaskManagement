package repository

import (
	"ManageTask/models"

	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) AddTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepo) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepo) GetOneTask(id int) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return task, err
}

func (r *TaskRepo) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepo) DeleteTask(id int) error {
	return r.db.Delete(&models.Task{}, id).Error
}
