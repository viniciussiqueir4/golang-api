package migrations

import (
	"example/webapi/models"

	"gorm.io/gorm"
)

func RunMIgrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
