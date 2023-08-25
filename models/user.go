package models

import "time"

type User struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	UserName    string `json:"user_name"`
	Password    string `json:"pass"`
	Role        string `json:"role"`
	IsActivated bool   `json:"is_activated"`
}
