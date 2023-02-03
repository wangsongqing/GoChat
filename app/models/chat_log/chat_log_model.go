package chat_log

import (
	"GoChat/app/models"
	"GoChat/pkg/database"
)

type ChatLog struct {
	models.BaseModel
	models.CommonTimestampsField

	TargetId int    `json:"target_id"`
	UserId   int    `json:"user_id"`
	Type     int    `json:"type"`
	Media    int    `json:"media"`
	Content  string `json:"content"`
}

func (cl *ChatLog) Create() int {
	if res := database.DB.Create(&cl); res.Error != nil {
		return 0
	}

	return cl.ID
}
