package controllers

import (
	"net/http"

	"example.com/go-back-end/config"
	"example.com/go-back-end/models"
	"github.com/gin-gonic/gin"
)

func BookCreate(c *gin.Context) {
	book := models.Book{
		Title:    "Book 1",
		AuthorID: 1,
	}
	result := config.DB.Create(&book)
	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "BookCreate",
		"book":    book,
	})

}
