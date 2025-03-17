package routes

import (
	"example.com/go-back-end/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	r.GET("/api/users", controllers.UserIndex)
	r.GET("/api/user/:id", controllers.UserShow)
	//send email
	r.POST("/api/send_email", controllers.SendEmail)
	r.POST("/api/create_user", controllers.UserCreate)
	r.PUT("/api/update_user/:id", controllers.UserUpdate)
	r.DELETE("/api/delete_user/:id", controllers.UserDelete)

	return r
}
