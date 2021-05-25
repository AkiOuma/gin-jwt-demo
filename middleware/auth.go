package middleware

import (
	"jwt-demo/controller"
	"jwt-demo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/login" && c.Request.Method == "POST" {
			c.Next()
		} else {
			token := c.GetHeader("Authorization")
			if token == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "operation is not authorized",
				})
				c.Abort()
				return
			}
			_, err := utils.ValidJwt(token[len("Bearer "):], controller.JwtKey)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				c.Abort()
				return
			}
			c.Next()
		}
	}
}
