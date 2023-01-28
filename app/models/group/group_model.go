package group

import (
	"GoChat/app/models"
	"GoChat/pkg/database"
)

type GroupChat struct {
	models.BaseModel
	models.CommonTimestampsField

	Name    string `json:"name"`
	OwnerId int    `json:"owner_id"`
	Type    int    `json:"type"`
	Desc    string `json:"desc"`
}

func (gc *GroupChat) Create() int64 {
	result := database.DB.Create(&gc)
	return result.RowsAffected
}
