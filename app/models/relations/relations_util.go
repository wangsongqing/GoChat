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
func GetRelationList(OwnerId int) (RelationsModel []Relations) {
	database.DB.Where("owner_id = ?", OwnerId).Find(&RelationsModel)
	return
}

func GetRelationsId(OwnerId int, TargetId int) bool {
	var count int64
	database.DB.Model(Relations{}).Where("owner_id = ? and target_id = ?", OwnerId, TargetId).Count(&count)
	return count > 0
}

func DelPeople(OwnerId int, TargetId int) int64 {
	res := database.DB.Where("owner_id = ? and target_id = ?", OwnerId, TargetId).Delete(&Relations{})
	return res.RowsAffected
}
