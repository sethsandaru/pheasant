package models

import (
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	ID uint64 `json:"-" gorm:"primaryKey"`
	HasUUID
	UserID      uint64         `json:"user_id" gorm:"index"`
	Title       string         `json:"title" gorm:"index"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	User User `gorm:"foreignKey:UserID"`
}

type EntityModel interface {
}

type entityModelDependencies struct{}

func GetEntityModel() EntityModel {
	return &entityModelDependencies{}
}
