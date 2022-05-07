package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pheasant-api/app/helper"
)

// DB is the database connection.
var DB *gorm.DB

// Initialize migrates and sets up the database.
func Initialize(withMigration bool) {
	dbUser := helper.GetEnv("DATABASE_USERNAME", "golang")
	dbPass := helper.GetEnv("DATABASE_PASSWORD", "golang")
	dbHost := helper.GetEnv("DATABASE_HOST", "localhost")
	dbPort := helper.GetEnv("DATABASE_PORT", "5432")
	dbName := helper.GetEnv("DATABASE_NAME", "go_test")
	//additional := "charset=utf8mb4&parseTime=True&loc=Local"

	// Assemble the connection string.
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Connect to the database.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not open database connection")
	}

	// Migrate the schemas
	if withMigration {
		db.AutoMigrate(&User{})
		db.AutoMigrate(&ForgotPasswordToken{})
		db.AutoMigrate(&Release{})
	}

	DB = db
}
