package models

import "time"

type Book struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	Author      string `json:"author"`
	Year        uint   `json:"year"`
	Title       string `json:"title"`
	Quantity    uint   `json:"quantity"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}