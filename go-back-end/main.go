package main

import (
	"example.com/go-back-end/config"
	"example.com/go-back-end/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDatabase()
}

func main() {
	r := gin.Default()

	routes.InitBookRoutes(r)
	routes.InitUserRoutes(r)

	r.Run()

}
