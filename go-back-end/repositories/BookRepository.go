package repositories

import (
	"example.com/go-back-end/config"
	"example.com/go-back-end/models"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) Create(Book models.Book) (models.Book, error) {
	result := config.DB.Create(&Book)

	return Book, result.Error
}
func (r *BookRepository) Update(Book models.Book) (models.Book, error) {
	result := config.DB.Save(&Book)

	return Book, result.Error
}

func (r *BookRepository) FindAll() ([]models.Book, error) {
	var Books []models.Book
	result := config.DB.Find(&Books)

	return Books, result.Error
}

func (r *BookRepository) FindById(id string) (models.Book, error) {
	var Book models.Book
	result := config.DB.First(&Book, id)

	return Book, result.Error
}

func (r *BookRepository) Delete(Book models.Book) error {
	result := config.DB.Delete(&Book)

	return result.Error
}
