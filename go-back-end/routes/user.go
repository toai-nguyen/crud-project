package routes

import (
	"example.com/go-back-end/controllers"
	"example.com/go-back-end/repositories"
	"example.com/go-back-end/services"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r.GET("/api/users", userController.Index)
	r.GET("/api/user/:id", userController.ShowById)
	// r.POST("/api/send_email", controllers.SendEmail)
	r.POST("/api/create_user", userController.Create)
	r.PUT("/api/update_user/:id", userController.Update)
	r.DELETE("/api/delete_user/:id", userController.Delete)

}
