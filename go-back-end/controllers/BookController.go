package controllers

import (
	"net/http"

	"example.com/go-back-end/dto/request"
	"example.com/go-back-end/dto/response"
	"example.com/go-back-end/services"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService *services.BookService
}

func NewBookController(bookService *services.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}
func (ctrt *BookController) Index(c *gin.Context) {
	books, err := ctrt.bookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var bookResponses []response.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, response.FromBook(book))
	}
	c.JSON(http.StatusOK, gin.H{
		"books": bookResponses,
	})
}
func (ctrt *BookController) Create(c *gin.Context) {
	var req request.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := ctrt.bookService.CreateBook(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book created",
		"book":    response.FromBook(book),
	})
}
func (ctrt *BookController) ShowById(c *gin.Context) {
	id := c.Param("id")
	book, err := ctrt.bookService.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": response.FromBook(book),
	})
}
func (ctrt *BookController) Update(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := ctrt.bookService.UpdateBook(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated",
		"book":    response.FromBook(book),
	})
}
func (ctrt *BookController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrt.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted",
	})
}
