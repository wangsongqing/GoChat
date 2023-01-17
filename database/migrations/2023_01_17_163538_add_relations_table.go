package migrations

import (
	"GoChat/app/models"
	"GoChat/pkg/migrate"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func init() {

	type Relations struct {
		models.BaseModel

		OwnerId  string `gorm:"type:int;not null;index;default:0;comment:用户ID" json:"owner_id"`
		TargetId string `gorm:"type:int;index;default:0;comment:好友ID" json:"target_id"`
		Type     string `gorm:"type:int;default:0;comment:好友类型  1私人好友  2群好友" json:"type"`
		Desc     string `gorm:"type:varchar(255);comment:描述" json:"desc"`

		models.CommonTimestampsField

		DeleteAt time.Time `gorm:"column:delete_at;" json:"deleted_at,omitempty"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Relations{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Relations{})
	}

	migrate.Add("2023_01_17_163538_add_relations_table", up, down)
}
