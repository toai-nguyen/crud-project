package main

import (
	"example.com/go-back-end/config"
	"example.com/go-back-end/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDatabase()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Book{})
}
