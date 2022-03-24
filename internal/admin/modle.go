package admin

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
type UpdateAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginAdmin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
