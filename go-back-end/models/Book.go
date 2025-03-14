package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID       string `json:"id"`
	Title    string `json:"title"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author" gorm:"foreignKey:AuthorID"`
}
