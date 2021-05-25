package controller

import (
	"jwt-demo/service"
	"jwt-demo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var JwtKey = []byte("gin-jwt")

type LoginInfo struct {
	Userid   string `json:"userid"`
	Password string `json:"password"`
}

func LoginController(app *gin.Engine) {
	app.POST("/api/login", AuthUser)
}

func AuthUser(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.ShouldBind(&loginInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := service.GetUser(loginInfo.Userid, loginInfo.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user authentication failed",
		})
		return
	}

	token := utils.SignJwt(JwtKey, user)

	c.JSON(http.StatusOK, gin.H{
		"message": "login succesfully",
		"data":    token,
	})
}
