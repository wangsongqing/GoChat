package v1

import (
	"GoChat/app/models/chat_log"
	"GoChat/pkg/auth"
	"GoChat/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChatLogController struct {
	BaseAPIController
}

func (cc *ChatLogController) List(c *gin.Context) {

	userInfo := auth.CurrentUser(c)
	getTargetId := c.Query("target_id")
	if getTargetId == "" {
		response.JSON(c, gin.H{
			"code": -1,
			"msg":  "参数target_id为必填",
		})
		return
	}

	targetId, _ := strconv.Atoi(getTargetId)
	data, pager := chat_log.Paginate(c, 2, userInfo.ID, targetId)
	response.JSON(c, gin.H{
		"code":  1,
		"msg":   "success",
		"data":  data,
		"pager": pager,
	})
}
