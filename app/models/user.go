package models

import (
	"pheasant-api/database"
	"time"
)

// User is the main user model.
type User struct {
	ID        int    `json:"id" gorm:"->;primaryKey"`
	Email     string `json:"email" gorm:"index:,unique"`
	Password  string
	FullName  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUserByEmail(email string) (*User, error) {
	user := User{}
	userResult := database.DB.Where("email = ?", email).First(&user)
	if userResult.Error != nil {
		return nil, userResult.Error
	}

	return &user, nil
}
