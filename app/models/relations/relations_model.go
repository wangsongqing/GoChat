package relations

import (
	"GoChat/app/models"
	"GoChat/pkg/database"
)

type Relations struct {
	models.BaseModel
	ID       int
	OwnerId  int    `json:"owner_id"`
	TargetId int    `json:"target_id"`
	Type     int    `json:"type"`
	Desc     string `json:"desc"`

	models.CommonTimestampsField
}

// Create 新增数据
func (r *Relations) Create() int64 {
	row := database.DB.Create(&r)
	return row.RowsAffected
}
