// Package routes 注册路由
package routes

import (
	controllers "GoChat/app/http/controllers/api/v1"
	"GoChat/app/http/controllers/api/v1/auth"
	"GoChat/app/http/middlewares"
	middlewareapp "GoChat/app/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 静态图片访问url
	r.StaticFS("/public/uploads", http.Dir("./public/uploads"))

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := auth.SignupController{}
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/email", suc.SignupUsingEmail)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)

			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)

			user := new(auth.UserController)
			authGroup.POST("/getUser", user.GetUser)

			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)

			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 使用邮箱登录
			authGroup.POST("/login/using-email", lgc.LoginByEmail)
			// 账号密码登录
			authGroup.POST("/login/account", lgc.LoginByPassword)
			// 刷新 token
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			// 手机修改密码
			pwc := auth.PasswordController{}
			authGroup.POST("/password-reset/using-phone", middlewares.LimitPerRoute("20-H"), pwc.PhoneByPassword)
			authGroup.POST("/password-reset/using-email", middlewares.LimitPerRoute("20-H"), pwc.EmailByPassword)
		}

		uc := controllers.UsersController{}
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("/user", middlewareapp.AuthJWT(), uc.CurrentUser)
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middlewareapp.AuthJWT(), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewareapp.AuthJWT(), uc.UpdateEmail)
			usersGroup.PUT("/phone", middlewareapp.AuthJWT(), uc.UpdatePhone)
			usersGroup.PUT("/update_password", middlewareapp.AuthJWT(), uc.UpdatePassword)
			usersGroup.PUT("/avatar", middlewareapp.AuthJWT(), uc.UpdateAvatar)
		}

		rc := controllers.RelationsController{}
		relationsGroup := v1.Group("/chat")
		{
			relationsGroup.GET("", middlewareapp.AuthJWT(), rc.List)
			relationsGroup.POST("/add", middlewareapp.AuthJWT(), rc.Add)
		}

		mc := controllers.MessageController{}
		messageGroup := v1.Group("msg")
		{
			messageGroup.GET("/send", mc.SendMsg)
		}

	}

}
