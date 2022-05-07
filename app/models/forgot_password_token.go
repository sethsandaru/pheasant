package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ForgotPasswordToken struct {
	ID        uint64         `json:"-" gorm:"primaryKey"`
	UserID    uint64         `json:"user_id" gorm:"index"`
	Token     string         `json:"description" gorm:"index"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	ExpiredAt time.Time      `json:"expired_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User User `gorm:"foreignKey:UserID"`
}

type ForgotPasswordTokenModel interface {
	IsTokenStillValid(token string) bool
	FindByToken(token string) *ForgotPasswordToken
	CreateForUser(user User) (*ForgotPasswordToken, error)
	DeleteByToken(token string) bool
}

type forgotPasswordTokenDependencies struct{}

func GetForgotPasswordTokenModel() ForgotPasswordTokenModel {
	return &forgotPasswordTokenDependencies{}
}

const tokenExpirationMinutes = 30

func (model *forgotPasswordTokenDependencies) CreateForUser(user User) (*ForgotPasswordToken, error) {
	now := time.Now()
	expiredAt := time.Now().Add(time.Minute * tokenExpirationMinutes)

	// delete all existing token
	DB.Where("user_id = ?", user.ID).Delete(&ForgotPasswordToken{})

	// generate new token
	forgotPasswordToken := &ForgotPasswordToken{
		UserID:    user.ID,
		Token:     uuid.NewString(),
		CreatedAt: now,
		ExpiredAt: expiredAt,
	}

	createResult := DB.Create(forgotPasswordToken)
	if createResult.Error != nil {
		return nil, createResult.Error
	}

	return forgotPasswordToken, nil
}

func (model *forgotPasswordTokenDependencies) IsTokenStillValid(token string) bool {
	now := time.Now()
	tokenQueryResult := DB.Where("token = ? AND expired_at > ?", token, now).First(&ForgotPasswordToken{})

	return tokenQueryResult.Error == nil
}

func (model *forgotPasswordTokenDependencies) FindByToken(token string) *ForgotPasswordToken {
	forgotPasswordToken := ForgotPasswordToken{}
	tokenQueryResult := DB.Preload("User").Where("token = ?", token).First(&forgotPasswordToken)
	if tokenQueryResult.Error != nil {
		return nil
	}

	return &forgotPasswordToken
}

func (model *forgotPasswordTokenDependencies) DeleteByToken(token string) bool {
	deleteResult := DB.Where("token = ?", token).Delete(&ForgotPasswordToken{})

	return deleteResult.Error != nil
}
