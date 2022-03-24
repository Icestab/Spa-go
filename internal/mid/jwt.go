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
		token := c.GetHeader("token")
		if token == "" {
			code = -1
		} else {
			claims, err := auth.ParseToken(token)
			log.Println(claims)
			if err != nil {
				code = -2
			}
		}

		if code != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "Auth Error",
				"data":    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
