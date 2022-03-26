package user

import (
	"errors"
	"spa-go/internal/utils/db"
)

var (
	ErrNotFound = errors.New("User not found.")

	ErrUsernameOrPhoneNil = errors.New("Username or phone can not be null.")
)

func Get(uid string) (*GetUserInfo, error) {
	var u = GetUserInfo{}

	err := db.Sqlite.Table("user").Where("id = ?", uid).Where("deleted_at is  null").First(&u).Error
	if err != nil {
		return nil, ErrNotFound
	}
	return &u, nil
}

func Create(cu CreateUser) error {
	u := User{Name: cu.Name, Phone: cu.Phone, Total: cu.Total, Used: cu.Used}
	err := db.Sqlite.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
func Delete(uid string) error {
	_, err := Get(uid)
	if err != nil {
		return err
	}
	err = db.Sqlite.Delete(&User{}, "id = ?", uid).Error
	if err != nil {
		return err
	}
	return nil
}
func Update(uid string, uu UpdateUser) error {
	u := User{Name: uu.Name, Phone: uu.Phone}
	_, err := Get(uid)
	if err != nil {
		return err
	}
	err = db.Sqlite.Model(&User{}).Where("id = ?", uid).Updates(u).Error
	if err != nil {
		return err
	}
	return nil
}
