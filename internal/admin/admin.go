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

	ErrUsernameOrPasswordNil = errors.New("用户名或密码不能为空")

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
