package auth

import (
	v1 "GoChat/app/http/controllers/api/v1"
	"GoChat/app/models/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	v1.BaseAPIController
}

func (sc *UserController) GetUser(c *gin.Context) {

	// 请求对象
	type PhoneExistRequest struct {
		Id int `json:"id"`
	}

	request := PhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": user.GetUser(request.Id),
	})
}
