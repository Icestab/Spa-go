package v1

import (
	"spa-go/cmd/app/handlers"
	"spa-go/internal/admin"

	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	var a admin.LoginAdmin
	if err := c.ShouldBindJSON(&a); err != nil {
		handlers.Re(c, -1, err.Error(), nil)
		return
	}
	token, err := admin.Token(a)
	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)
	} else {
		t := map[string]string{"token": token}
		handlers.Re(c, 0, "success", t)
	}
}
