package controller

import (
	"jwt-demo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(app *gin.Engine) {
	app.GET("/api/user", GetAllUser)
}

func GetAllUser(c *gin.Context) {
	users := service.GetAllUser()
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user successfully",
		"data":    users,
	})
}
