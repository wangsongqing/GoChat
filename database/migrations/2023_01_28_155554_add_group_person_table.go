package migrations

import (
	"GoChat/app/models"
	"GoChat/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type GroupPeople struct {
		models.BaseModel

		GroupId int `gorm:"default:0;not null;index"`
		UserId  int `gorm:"default:0;index;not null"`
		Status  int `gorm:"type:tinyint(1);default:0;not null;comment:1:正常 2禁言 3退群"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&GroupPeople{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&GroupPeople{})
	}

	migrate.Add("2023_01_28_155554_add_group_person_table", up, down)
}
