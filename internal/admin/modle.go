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
	Username string `json:"username" `
	Password string `json:"password" `
}
type UpdatePassword struct {
	OldPassword string `json:"oldPassword" `
	NewPassword string `json:"newPassword" `
}
