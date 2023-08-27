package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	BookID  uint `json:"book_id"`
}