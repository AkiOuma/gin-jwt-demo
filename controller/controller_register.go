package controller

import "github.com/gin-gonic/gin"

func ControllerRegister(app *gin.Engine) {
	LoginController(app)
	UserController(app)
}
