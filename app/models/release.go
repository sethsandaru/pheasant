package models

import (
	"time"
)

type Release struct {
	ID          int       `json:"id" gorm:"->;primaryKey"`
	Version     string    `json:"version" gorm:"index"`
	Description string    `json:"description" gorm:"index"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
