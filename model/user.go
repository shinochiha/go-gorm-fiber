package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email" validate:"required,email"`
	Password  string         `json:"-" gorm:"column:password" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
