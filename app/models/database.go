package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"pheasant-api/app/helper"
	"strconv"
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
	// TODO: change to goose if possible
	if withMigration {
		db.AutoMigrate(&User{})
		db.AutoMigrate(&ForgotPasswordToken{})
		db.AutoMigrate(&Release{})
		db.AutoMigrate(&Entity{})
	}

	DB = db
}

type HasUUID struct {
	UUID string `json:"uuid" gorm:"index:,unique; default: uuid_generate_v4()"`
}

func findByUuidQuery(uuid string) *gorm.DB {
	return DB.Where("uuid = ?", uuid)
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
