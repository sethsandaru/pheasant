package models

import (
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
