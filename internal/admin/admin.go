package admin

import (
	"spa-go/internal/auth"
	"spa-go/internal/utils/db"
	"spa-go/internal/utils/sha256"

	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("用户不存在")

	ErrIncorrectUsernameOrPassword = errors.New("用户或密码错误")
	ErrIncorrectOldPassword        = errors.New("旧密码错误")
	ErrOldPasswordOrNewPasswordNil = errors.New("新密码或旧密码不能为空")
	ErrUsernameOrPasswordNil       = errors.New("用户名或密码不能为空")

	ErrGenerateTokenFailed = errors.New("生成token失败")
)

func Token(pl LoginAdmin) (string, error) {
	if pl.Username == "" || pl.Password == "" {
		return "", ErrUsernameOrPasswordNil
	}
	pl.Password = sha256.Sha256(pl.Password)
	var admin Admin
	err := db.Sqlite.Where("username = ?", pl.Username).Where("password = ?", pl.Password).First(&admin).Error

	if err != nil {
		return "", ErrIncorrectUsernameOrPassword
	}

	token, err := auth.GenerateToken(admin.ID)
	if err != nil {
		return "", ErrGenerateTokenFailed
	}
	return token, nil
}
func UpdatePasswd(pl UpdatePassword) error {
	if pl.OldPassword == "" || pl.NewPassword == "" {
		return ErrOldPasswordOrNewPasswordNil
	}
	pl.OldPassword = sha256.Sha256(pl.OldPassword)
	pl.NewPassword = sha256.Sha256(pl.NewPassword)
	var admin Admin
	err := db.Sqlite.Where("password = ?", pl.OldPassword).First(&admin).Error
	if err != nil {
		return ErrIncorrectOldPassword
	}
	err = db.Sqlite.Model(&Admin{}).Where("id = ?", admin.ID).Updates(map[string]interface{}{"password": pl.NewPassword}).Error
	if err != nil {
		return err
	}
	return nil
}
