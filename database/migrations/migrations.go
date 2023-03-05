package migrations

import (
	models "github.com/lucasdk3/maui-oqcomer-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
