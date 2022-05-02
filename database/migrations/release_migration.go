package migrations

import (
	"pheasant-api/app/models"
	"pheasant-api/database"
)

func MigrateRelease() {
	database.DB.AutoMigrate(&models.Release{})
}
