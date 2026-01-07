package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ManageTask/models"
)

var DB *gorm.DB

func Connect() error {
	dsn := "postgres://urantune:Seigakartisde9@localhost:5432/TaskManagement?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		return err
	}
	DB = db
	return nil
}
