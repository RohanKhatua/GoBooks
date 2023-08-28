package models

import "time"

type Purchase struct {
	ID           uint `gorm:"primaryKey"`
	UserID       int `json:"user_id"`
	BookID       uint `json:"book_id"`
	Quantity     uint  `json:"quantity"`
	User         User `gorm:"foreignKey:UserID"`
	Book         Book `gorm:"foreignKey:BookID"`
	PurchaseDate time.Time
}
