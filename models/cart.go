package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ID 	uint `gorm:"primaryKey"`
	UserID  uint `json:"user_id" gorm:"index;constraint:OnDelete:CASCADE;references:User"`
    BookID  uint `json:"book_id" gorm:"index;constraint:OnDelete:CASCADE;references:Book"`
}