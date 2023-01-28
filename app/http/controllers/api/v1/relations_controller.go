package v1

import (
	"GoChat/app/models/relations"
	"GoChat/app/models/user"
	"GoChat/pkg/auth"
	"GoChat/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RelationsController struct {
	BaseAPIController
}

// List 获取好友列表
func (r *RelationsController) List(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	data := relations.GetRelationList(userModel.ID)
	response.JSON(c, gin.H{
		"data": data,
	})
}

// Add 添加好友
func (r *RelationsController) Add(c *gin.Context) {
	typeName := c.PostForm("type")
	desc := c.PostForm("desc")
	TargetName := c.PostForm("target_name")
	tarUser := user.GetByName(TargetName)
	if tarUser.ID == 0 {
		response.Abort500(c, "您添加的好友不存在，请稍后尝试~")
		return
	}

	userModel := auth.CurrentUser(c)

	if userModel.ID == tarUser.ID {
		response.Abort500(c, "自己不可以添加自己为好友哦，请稍后尝试~")
		return
	}
	relation := relations.Relations{}
	relation.OwnerId = userModel.ID
	relation.TargetId = tarUser.ID
	relation.Type, _ = strconv.Atoi(typeName)
	relation.Desc = desc

	if row := relation.Create(); row == 0 {
		response.Abort500(c, "写入失败，请稍后尝试~")
		return
	}

	response.JSON(c, gin.H{
		"code": 1,
		"msg":  "添加成功",
	})
}
