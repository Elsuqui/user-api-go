package migrations

import (
	"UserRestApi/models"

	"gorm.io/gorm"
)

func Boot(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
