package relations

import (
	"GoChat/pkg/app"
	"GoChat/pkg/database"
	"GoChat/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []Relations, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Relations{}),
		&users,
		app.V1URL(database.TableName(&Relations{})),
		perPage,
	)
	return
}

// GetRelationList 获取好友列表
func GetRelationList(OwnerId int) (RelationsMolde []Relations) {
	database.DB.Where("owner_id = ?", OwnerId).Find(&RelationsMolde)
	return
}
