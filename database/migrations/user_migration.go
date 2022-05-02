package migrations

import (
	"pheasant-api/app/models"
	"pheasant-api/database"
)

func MigrateUser() {
	database.DB.AutoMigrate(&models.User{})
}
