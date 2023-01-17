// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "GoChat/app/http/controllers/api/v1"
	"GoChat/app/models/user"
	"GoChat/app/requests"
	"GoChat/pkg/hash"
	_ "GoChat/pkg/helpers"
	"GoChat/pkg/jwt"
	"GoChat/pkg/response"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist)
	if !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"name":  request.Name,
		"phone": request.Phone,
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检查邮箱是否有注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

// SignupUsingEmail 邮件注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	request := requests.SignupUsingEmail{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmail); !ok {
		return
	}

	_data := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hash.BcryptHash(request.Password),
	}
	_data.Create()
	if _data.ID > 0 {
		token := jwt.NewJWT().IssueToken(_data.GetStringID(), _data.Name)
		response.JSON(c, gin.H{
			"token": token,
			"data":  _data,
		})
	} else {
		response.JSON(c, gin.H{
			"ok": false,
		})
	}

}

//SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}

	// 1. 验证表单
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 验证成功，创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: hash.BcryptHash(request.Password),
	}

	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
