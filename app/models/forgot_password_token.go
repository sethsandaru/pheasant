package models

import (
	"time"
)

type ForgotPasswordToken struct {
	ID        uint64    `json:"-" gorm:"primaryKey"`
	UserID    string    `json:"user_id" gorm:"index"`
	Token     string    `json:"description" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
	ExpiredAt time.Time `json:"expired_at" gorm:"index"`

	User User `gorm:"foreignKey:UserID"`
}
