package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	UserName    string `json:"user_name"`
	Password    string `json:"pass"`
	Role        string `json:"role"`
	IsActivated bool   `json:"is_activated"`
	Reviews []Review //Many Reviews
	CartItems []CartItem //Many Cart Items
}
