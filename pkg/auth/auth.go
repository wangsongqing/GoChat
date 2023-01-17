package auth

import (
	"GoChat/app/models/user"
	"GoChat/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
)

func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机未注册")
	}

	return userModel, nil
}

func LoginByEmail(email string) (user.User, error) {
	userModel := user.GetByEmail(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("邮箱未注册")
	}

	return userModel, nil
}

func LoginByAccount(account string) (user.User, error) {
	userModel := user.GetAccount(account)
	if userModel.ID != 0 {
		return userModel, nil
	}

	return user.User{}, errors.New("账户不存在")
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	UserModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return UserModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
