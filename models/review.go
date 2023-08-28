package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    uint    `json:"user_id" gorm:"index;constraint:OnDelete:CASCADE;references:User"`
    BookID    uint    `json:"book_id" gorm:"index;constraint:OnDelete:CASCADE;references:Book"`
	ReviewDate time.Time
	Rating    int    `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	Comment    string `json:"comment"`
	User       User   `gorm:"foreignKey:UserID"`
	Book       Book   `gorm:"foreginKey:BookID"`
}
