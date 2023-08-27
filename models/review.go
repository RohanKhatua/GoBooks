package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
	ReviewDate time.Time
	Rating int `json:"rating"`
	Comment string `json:"comment"`
	User User
	Book Book
}