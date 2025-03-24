package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string `json:"title"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
	Status   string `json:"status"`
	AuthorID uint   `json:"author_id"`
}
