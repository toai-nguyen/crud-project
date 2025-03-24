package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Role     string  `json:"role"`
	Password string  `json:"password"`
	Books    []Book  `gorm:"foreignKey:AuthorID"`
}
