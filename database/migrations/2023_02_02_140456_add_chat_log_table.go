package migrations

import (
	"GoChat/app/models"
	"GoChat/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ChatLog struct {
		models.BaseModel

		TargetId int    `gorm:"type:int(11);not null;index;default:0;comment:信息接收者ID,如果是群则为群ID"`
		UserId   int    `gorm:"type:int(11);not null;index;default:0;comment:信息发送者ID"`
		Type     int    `gorm:"type:tinyint(1);not null;default:0;comment:消息的类型: 1私聊  2群聊"`
		Media    int    `gorm:"type:tinyint(1);not null;default:0;comment:信息类型:1文字"`
		Content  string `gorm:"type:text;comment:消息内容"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ChatLog{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ChatLog{})
	}

	migrate.Add("2023_02_02_140456_add_chat_log_table", up, down)
}
