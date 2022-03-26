package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Total int    `json:"total"`
	Used  int    `json:"used"`
}
type CreateUser struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Total int    `json:"total"`
	Used  int    `json:"used"`
}
type UpdateUser struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
type GetUserInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Total int    `json:"total"`
	Used  int    `json:"used"`
}
