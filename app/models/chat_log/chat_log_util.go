package chat_log

import (
	"GoChat/pkg/app"
	"GoChat/pkg/database"
	"GoChat/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int, userId int, targetId int, typeId int) (chatLog []ChatLog, paging paginator.Paging) {
	paging = paginator.PaginateChatLog(
		c,
		database.DB.Model(ChatLog{}),
		&chatLog,
		app.V1URL(database.TableName(&ChatLog{})),
		perPage,
		userId,
		targetId,
		typeId,
	)
	return
}
