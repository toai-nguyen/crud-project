package response

import "example.com/go-back-end/models"

type BookResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
	Status   string `json:"status"`
	AuthorID uint   `json:"author_id"`
}

func FromBook(book models.Book) BookResponse {
	return BookResponse{
		ID:       book.ID,
		Title:    book.Title,
		Price:    book.Price,
		Quantity: book.Quantity,
		Status:   book.Status,
		AuthorID: book.AuthorID,
	}
}
