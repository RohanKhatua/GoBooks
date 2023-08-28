package models

import "time"

type Purchase struct {
	ID           uint `gorm:"primaryKey"`
	UserID       int   `json:"user_id" gorm:"index;constraint:OnDelete:CASCADE;references:User"`
    BookID       uint  `json:"book_id" gorm:"index;constraint:OnDelete:CASCADE;references:Book"`
	Quantity     uint `json:"quantity"`
	User         User `gorm:"foreignKey:UserID"`
	Book         Book `gorm:"foreignKey:BookID"`
	PurchaseDate time.Time
}
