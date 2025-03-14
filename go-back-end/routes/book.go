package routes

import (
	"example.com/go-back-end/controllers"
	"github.com/gin-gonic/gin"
)

func InitBookRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/book", controllers.BookCreate)

	return r
}
