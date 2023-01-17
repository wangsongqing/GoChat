package migrations

import (
	"GoChat/app/models"
	"GoChat/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		Name         string `gorm:"type:varchar(255);not null;index"`
		Email        string `gorm:"type:varchar(255);index;default:null"`
		Phone        string `gorm:"type:varchar(20);index;default:null"`
		Password     string `gorm:"type:varchar(255)"`
		City         string `gorm:"type:varchar(100);not null;default:''"`
		Introduction string `gorm:"type:varchar(255);not null;default:''"`
		Avatar       string `gorm:"type:varchar(255);not null;default:''"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_09_19_164409_add_users_table", up, down)
}
