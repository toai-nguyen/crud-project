package controllers

import (
	"net/http"
	"net/smtp"
	"os"

	"example.com/go-back-end/dto/request"
	"example.com/go-back-end/dto/response"
	"example.com/go-back-end/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ctrl *UserController) Index(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.FromUser(user))
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userResponses,
	})
}

func (ctrl *UserController) Create(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"user":    response.FromUser(user),
	})
}

func (ctrl *UserController) ShowById(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": response.FromUser(user),
	})
}

func (ctrl *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.UpdateUser(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
		"user":    response.FromUser(user),
	})
}
func (ctrl *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
func SendEmail(c *gin.Context) {
	var RequestBody struct {
		Subject  string `json:"subject" binding:"required"`
		Body     string `json:"body" binding:"required"`
		Receiver string `json:"receiver" binding:"required"`
	}
	from := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	// host := os.Getenv("EMAIL_HOST")

	c.Bind(&RequestBody)

	toEmailAddress := RequestBody.Receiver
	to := []string{toEmailAddress}

	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_POST")

	address := host + ":" + port

	headers := "From: " + from + "\r\n" +
		"To: " + RequestBody.Receiver + "\r\n" +
		"Subject: " + RequestBody.Subject + "\r\n"
	mess := []byte(headers + RequestBody.Body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, mess)

	if err != nil {
		c.JSON(400, gin.H{
			"messsage": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "sending to email",
	})

}
