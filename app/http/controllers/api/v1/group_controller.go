package v1

import (
	"GoChat/app/models/group"
	"GoChat/app/models/group_people"
	"GoChat/app/models/user"
	"GoChat/pkg/auth"
	"GoChat/pkg/response"
	"github.com/gin-gonic/gin"
)

type GroupController struct {
	BaseAPIController
}

// CreateGroup 新建群
func (gc *GroupController) CreateGroup(c *gin.Context) {
	groupChat := group.GroupChat{}
	c.BindJSON(&groupChat)

	if len(groupChat.Name) == 0 {
		response.Abort500(c, "群名称不能为空")
		return
	}

	userModel := auth.CurrentUser(c)
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

// Add 添加用户进群
func (gc *GroupController) Add(c *gin.Context) {
	gp := group_people.GroupPeople{}
	c.BindJSON(&gp)

	if gp.GroupId == 0 || gp.UserId == 0 {
		response.Abort500(c, "group_id和user_id为必填")
		return
	}

	if isUser := user.IsUserExist(gp.UserId); isUser == false {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "用户ID不存在",
		})
		return
	}

	if ok := group_people.IsExistGroup(gp.UserId, gp.GroupId); ok == true {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "你已经在群里了",
		})
		return
	}

	gp.Status = 1
	gp.Create()

	response.JSON(c, gin.H{
		"code": 1,
		"msg":  "success",
	})
}

// PopGroupPeople 用户退群
func (gc *GroupController) PopGroupPeople(c *gin.Context) {
	gp := group_people.GroupPeople{}
	c.BindJSON(&gp)

	if gp.GroupId == 0 {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "group_id能不能为空",
		})
		return
	}

	userModel := auth.CurrentUser(c)
	if userModel.ID == 0 {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "用户错误",
		})
		return
	}
	//status := map[string]int{"status": 3}
	if result := gp.UpdateStatus(gp.GroupId, userModel.ID); result == 0 {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "退群失败",
		})
		return
	}

	response.JSON(c, gin.H{
		"code": 1,
		"msg":  "success",
	})
}
