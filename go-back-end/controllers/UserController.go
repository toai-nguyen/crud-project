package controllers

import (
	"net/smtp"
	"os"

	"example.com/go-back-end/config"
	"example.com/go-back-end/models"
	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}
func UserCreate(c *gin.Context) {

	var RequestBody struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}
	c.Bind(&RequestBody)
	user := models.User{
		Name:  RequestBody.Name,
		Email: RequestBody.Email,
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"message": "User created",
		"user":    user,
	})
}
func UserShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	config.DB.First(&user, id)
	c.JSON(200, gin.H{
		"user": user,
	})
}
func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	config.DB.First(&user, id)
	var RequestBody struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}
	c.Bind(&RequestBody)
	user.Name = RequestBody.Name
	user.Email = RequestBody.Email

	config.DB.Save(&user)
	c.JSON(200, gin.H{
		"message": "User updated",
		"user":    user,
	})
}
func UserDelete(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	config.DB.First(&user, id)
	config.DB.Delete(&user)
	c.JSON(200, gin.H{
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

	host := os.Getenv("")
	port := "587"

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
