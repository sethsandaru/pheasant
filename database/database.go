package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pheasant-api/app/helper"
	"pheasant-api/database/migrations"
)

// DB is the database connection.
var DB *gorm.DB

// Initialize migrates and sets up the database.
func Initialize() {
	u := helper.GetEnv("DATABASE_USER", "golang")
	p := helper.GetEnv("DATABSE_PASSWORD", "golang")
	h := helper.GetEnv("DATABASE_HOST", "localhost:3306")
	n := helper.GetEnv("DATABASE_NAME", "go_test")
	q := "charset=utf8mb4&parseTime=True&loc=Local"

	// Assemble the connection string.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", u, p, h, n, q)

	// Connect to the database.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the schemas
	migrations.MigrateUser()
	migrations.MigrateRelease()

	if err != nil {
		panic("Could not open database connection")
	}

	DB = db
}