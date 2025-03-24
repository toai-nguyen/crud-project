package routes

import (
	"example.com/go-back-end/controllers"
	"example.com/go-back-end/repositories"
	"example.com/go-back-end/services"
	"github.com/gin-gonic/gin"
)

func InitBookRoutes(r *gin.Engine) {
	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)

	r.GET("api/books", bookController.Index)
	r.GET("api/book/:id", bookController.ShowById)
	r.PUT("api/book_update/:id", bookController.Update)
	r.DELETE("api/book_delete/:id", bookController.Delete)
	r.POST("api/book_create", bookController.Create)
}
