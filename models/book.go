package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Author      string `json:"author"`
	Year        uint   `json:"year"`
	Title       string `json:"title"`
	Quantity    uint   `json:"quantity"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Reviews []Review
}