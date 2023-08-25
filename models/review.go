package models

import "time"

type Review struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
	ReviewDate time.Time
	Rating int `json:"rating"`
	Comment string `json:"comment"`
}