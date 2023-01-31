package v1

import (
	"GoChat/app/models/relations"
	"GoChat/app/models/user"
	"GoChat/app/requests"
	"GoChat/pkg/auth"
	"GoChat/pkg/response"
	"github.com/gin-gonic/gin"
)

type RelationsController struct {
	BaseAPIController
}

// List 获取好友列表
func (rc *RelationsController) List(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	data := relations.GetRelationList(userModel.ID)
	response.JSON(c, gin.H{
		"data": data,
	})
}

// Add 添加好友
func (rc *RelationsController) Add(c *gin.Context) {

	relationRequest := requests.RelationDelRequest{}
	if ok := requests.Validate(c, &relationRequest, requests.CheckRelation); !ok {
		return
	}

	tarUser := user.GetByName(relationRequest.TargetName)
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
	relation.Type = relationRequest.Type
	relation.Desc = relationRequest.Desc

	if row := relation.Create(); row == 0 {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "添加好友失败，请稍候再试~",
		})
	}

	response.JSON(c, gin.H{
		"code": 1,
		"msg":  "添加成功",
	})
}

// DelRelations 删除好友
func (rc *RelationsController) DelRelations(c *gin.Context) {
	relation := relations.Relations{}
	c.BindJSON(&relation)

	if relation.TargetId == 0 {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "target_id好友ID为必填",
		})
		return
	}

	userModel := auth.CurrentUser(c)

	if userModel.ID == 0 {
		response.Abort500(c, "用户异常，请稍后尝试~")
		return
	}

	if res := relations.GetRelationsId(userModel.ID, relation.TargetId); res == false {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "该用户不是你的好友，删除好友失败~",
		})
		return
	}

	if delRes := relations.DelPeople(userModel.ID, relation.TargetId); delRes <= 0 {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "删除失败，请稍候再试~",
		})
		return
	}

	response.JSON(c, gin.H{
		"code": 1,
		"msg":  "success",
	})
}
