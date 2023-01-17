package requests

import (
	"GoChat/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginPasswordRequest struct {
	Account  string `json:"account,omitempty" valid:"account"`
	Password string `json:"password,omitempty" valid:"password"`
}

func LoginPassword(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"account":  []string{"required"},
		"password": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"account": []string{
			"required:登录账号必填，参数名称 account",
		},
		"password": []string{
			"required:登录密码必填",
			"digits:密码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*LoginPasswordRequest)
	errs = validators.ValidatePassword(_data.Account, _data.Password, errs)

	return errs
}
