package v1

import (
	"GoChat/app/models/group"
	"GoChat/app/models/group_people"
	"GoChat/pkg/auth"
	"GoChat/pkg/response"
	"github.com/gin-gonic/gin"
)

type GroupController struct {
	BaseAPIController
}

// CreateGroup 新建群
func (gc *GroupController) CreateGroup(c *gin.Context) {
	checkGroup := group.GroupChat{}
	c.BindJSON(&checkGroup)

	if len(checkGroup.Name) == 0 {
		response.Abort500(c, "群名称不能为空")
		return
	}

	userModel := auth.CurrentUser(c)
	groupChat := group.GroupChat{}
	groupChat.Name = checkGroup.Name
	groupChat.OwnerId = userModel.ID
	res := groupChat.Create()

	if res == 0 {
		response.Abort500(c, "写入失败，请稍后尝试~")
		return
	}

	response.JSON(c, gin.H{
		"code": 1,
		"data": "success",
	})
}

func (gc *GroupController) Add(c *gin.Context) {
	gp := group_people.GroupPeople{}
	c.BindJSON(&gp)

	if gp.GroupId == 0 || gp.UserId == 0 {
		response.Abort500(c, "group_id和user_id为必填")
		return
	}

	gp.Status = 1
	gp.Create()

	response.JSON(c, gin.H{
		"code": 1,
		"msg":  "success",
	})
}
