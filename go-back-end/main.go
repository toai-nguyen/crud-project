package main

import (
	"example.com/go-back-end/config"
	"example.com/go-back-end/routes"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDatabase()
}

func main() {
	userRoutes := routes.InitUserRoutes()
	bookRoutes := routes.InitBookRoutes()
	userRoutes.Run()
	bookRoutes.Run()
}
