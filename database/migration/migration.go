package migration

import (
	"projectapi/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.User{},
	)

	return err
}
