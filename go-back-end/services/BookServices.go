package services

import (
	"example.com/go-back-end/dto/request"
	"example.com/go-back-end/models"
	"example.com/go-back-end/repositories"
)

type BookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo *repositories.BookRepository) *BookService {
	return &BookService{bookRepo: *bookRepo}
}

func (s *BookService) CreateBook(req request.CreateBookRequest) (models.Book, error) {
	book := models.Book{
		Title:    req.Title,
		Price:    req.Price,
		Quantity: req.Quantity,
		Status:   req.Status,
		AuthorID: req.AuthorID,
	}
	return s.bookRepo.Create(book)
}
func (s *BookService) UpdateBook(id string, req request.UpdateBookRequest) (models.Book, error) {
	book, err := s.bookRepo.FindById(id)
	if err != nil {
		return book, err
	}

	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Price != 0 {
		book.Price = req.Price
	}
	if req.Quantity != 0 {
		book.Quantity = req.Quantity
	}
	if req.Status != "" {
		book.Status = req.Status
	}

	return s.bookRepo.Update(book)
}

func (s *BookService) DeleteBook(id string) error {
	book, err := s.bookRepo.FindById(id)
	if err != nil {
		return err
	}

	return s.bookRepo.Delete(book)
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.bookRepo.FindAll()
}
func (s *BookService) GetBookById(id string) (models.Book, error) {
	return s.bookRepo.FindById(id)
}
