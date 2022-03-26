package v1

import (
	"spa-go/cmd/app/handlers"
	"spa-go/internal/user"

	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	id := c.Param("id")
	u, err := user.Get(id)
	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)
	} else {
		handlers.Re(c, 0, "success", u)
	}
}
func UserCreate(c *gin.Context) {
	var cu user.CreateUser
	if err := c.ShouldBindJSON(&cu); err != nil {
		handlers.Re(c, -1, err.Error(), nil)
		return
	}

	err := user.Create(cu)
	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)
	} else {
		handlers.Re(c, 0, "success", nil)
	}
}
func UserDelete(c *gin.Context) {
	id := c.Param("id")
	err := user.Delete(id)
	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)
	} else {
		handlers.Re(c, 0, "success", nil)
	}
}
func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var uu user.UpdateUser
	if err := c.ShouldBindJSON(&uu); err != nil {
		handlers.Re(c, -1, err.Error(), nil)
		return
	}
	err := user.Update(id, uu)
	if err != nil {
		handlers.Re(c, -1, err.Error(), nil)
	} else {
		handlers.Re(c, 0, "success", nil)
	}
}
