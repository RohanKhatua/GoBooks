package models

import "time"

type Purchase struct {
	UserID       uint `json:"user_id"`
	BookID       uint `json:"book_id"`
	PurchaseDate time.Time
}