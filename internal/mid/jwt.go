package middleware

import (
	"log"
	"net/http"

	"spa-go/internal/auth"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 0
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 401
		} else {
			claims, err := auth.ParseToken(token)
			log.Println(claims)
			if err != nil {
				code = 403
			}
		}

		if code != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "Auth Error",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
