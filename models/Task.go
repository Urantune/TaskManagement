package models

type Task struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json:"name"`
	Status string `json:"status"`

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`
}
