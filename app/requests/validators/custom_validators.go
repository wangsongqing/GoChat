// Package validators 存放自定义规则和验证器
package validators

import (
	"GoChat/app/models/user"
	"GoChat/pkg/captcha"
	"GoChat/pkg/config"
	"GoChat/pkg/hash"
	"GoChat/pkg/verifycode"
)

// ValidateCaptcha 自定义规则，验证『图片验证码』
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

// ValidatePasswordConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	if config.Env("APP_ENV") != "local" {
		if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
			errs["verify_code"] = append(errs["verify_code"], "验证码错误")
		}
	}
	return errs
}

// ValidatePassword 账号密码验证
func ValidatePassword(account string, pwd string, errs map[string][]string) map[string][]string {
	userInfo := user.GetAccount(account)
	if userInfo.ID <= 0 {
		errs["password"] = append(errs["password"], "账号不存在")
	}

	if checkPwd := hash.BcryptCheck(pwd, userInfo.Password); !checkPwd {
		errs["password"] = append(errs["password"], "密码错误，请重试")
	}

	return errs
}

// ValidateCheckPassword 修改密码手机验证码
func ValidateCheckPassword(phone string, code string, errs map[string][]string) map[string][]string {
	userInfo := user.GetByPhone(phone)
	if userInfo.ID == 0 {
		errs["phone"] = append(errs["phone"], "账户不存在")
	}

	if config.Env("APP_ENV") != "local" {
		if ok := verifycode.NewVerifyCode().CheckAnswer(phone, code); !ok {
			errs["verify_code"] = append(errs["verify_code"], "验证码错误")
		}
	}

	return errs
}

// ValidateCheckPasswordEmail ValidateCheckPassword 修改密码邮箱验证码
func ValidateCheckPasswordEmail(email string, code string, errs map[string][]string) map[string][]string {
	userInfo := user.GetByEmail(email)
	if userInfo.ID == 0 {
		errs["email"] = append(errs["email"], "账户不存在")
	}

	if config.Env("APP_ENV") != "local" {
		if ok := verifycode.NewVerifyCode().CheckAnswer(email, code); !ok {
			errs["verify_code"] = append(errs["verify_code"], "验证码错误")
		}
	}

	return errs
}
