package admin

import (
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("用户不存在")

	ErrIncorrectUsernameOrPassword = errors.New("用户或密码错误")

	ErrUsernameOrPasswordNil = errors.New("用户名或密码不能为空")

	ErrGenerateTokenFailed = errors.New("生成token失败")
)
