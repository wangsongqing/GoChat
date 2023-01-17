package requests

import (
	"GoChat/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PasswordFindPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

type PasswordFindEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func PasswordFindEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{
			"required",
			"min:4",
			"max:20",
			"email",
		},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email为必填项，参数名称 email",
			"min:最少为4位字符",
			"max:最多能有20个字符",
			"email:必须为邮箱",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"password": []string{
			"required:密码必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*PasswordFindEmailRequest)
	errs = validators.ValidateCheckPasswordEmail(_data.Email, _data.VerifyCode, errs)
	return errs
}

func PasswordFindPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"password": []string{
			"required:密码必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*PasswordFindPhoneRequest)
	errs = validators.ValidateCheckPassword(_data.Phone, _data.VerifyCode, errs)
	return errs
}
