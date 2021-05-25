package main

import (
	"jwt-demo/controller"
	"jwt-demo/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(middleware.Auth())
	controller.ControllerRegister(server)
	server.Run(":8080")
}
