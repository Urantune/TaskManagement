package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"not null;uniqueIndex" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
	Role         string `gorm:"not null;default:user" json:"role"`
}
