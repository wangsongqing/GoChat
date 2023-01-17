package requests

import (
	"GoChat/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
	Name  string `json:"name,omitempty" valid:"name"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

type SignupUsingEmail struct {
	Name            string `json:"name,omitempty" valid:"name"`
	Email           string `json:"email,omitempty" valid:"email"`
	Code            string `json:"code,omitempty" valid:"code"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则ValidateSignupEmailExist
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	message := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return validate(data, rules, message)
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required: Email为必填项",
			"min: 长度需大于 4",
			"max: 长度小于 30",
			"email: Email格式不正确",
		},
	}

	return validate(data, rules, messages)
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

func ValidateSignupEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":    []string{"required", "min:4", "max:30", "email"},
		"code":     []string{"required", "len:6", "numeric"},
		"password": []string{"required", "min:6", "max:16"},
	}

	message := govalidator.MapData{
		"email": []string{
			"required: Email为必填项",
			"min: 长度需大于 4",
			"max: 长度小于 30",
			"email: Email格式不正确",
		},
		"code": []string{
			"required: 验证码必填",
			"len: 验证码必须是6位数字",
			"numeric: 验证码必须为数字",
		},
		"password": []string{
			"required: 密码必填",
			"min: 密码长度需大于 6",
			"max: 长度小于 16",
		},
	}

	errs := validate(data, rules, message)
	_data := data.(*SignupUsingEmail)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.Code, errs)

	return errs
}
