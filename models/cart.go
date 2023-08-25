package models

type Cart struct {
	UserID uint   `json:"user_id"`
	Books  []Book `json:"books"`
}