package group_people

import (
	"GoChat/app/models"
	"GoChat/pkg/database"
)

// StatusOk 用户群聊状态
const StatusOk = 1 // 正常

type GroupPeople struct {
	models.BaseModel
	models.CommonTimestampsField

	GroupId int `json:"group_id"`
	UserId  int `json:"user_id"`
	Status  int `json:"status"`
}

func (g *GroupPeople) Create() int64 {
	res := database.DB.Create(&g)
	return res.RowsAffected
}

func (g *GroupPeople) UpdateStatus(GroupId int, UserId int) int64 {
	res := database.DB.Model(GroupPeople{}).Where("group_id = ? and user_id = ?", GroupId, UserId).Update("status", 3)
	return res.RowsAffected
}
