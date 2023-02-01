package migrations

import (
	"GoChat/app/models"
	"GoChat/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type GroupChat struct {
		models.BaseModel

		Name    string `gorm:"type:varchar(255);not null;index;default:''"`
		OwnerId int    `gorm:"type:int(11);index;default:0;not null"`
		Type    int    `gorm:"type:int(11);not null;default:0"`
		Desc    string `gorm:"type:varchar(255);not null;default:''"`
		Status  int    `gorm:"type:int(11);not null;default:0"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&GroupChat{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&GroupChat{})
	}

	migrate.Add("2023_01_28_151059_add_group_chat_table", up, down)
}
