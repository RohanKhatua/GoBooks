package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	UserName    string `json:"user_name" gorm:"unique"`
	Password    string `json:"pass"`
	Role        string `json:"role"`
	IsActivated bool   `json:"is_activated" gorm:"default:true"`
}
