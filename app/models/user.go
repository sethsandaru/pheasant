package models

import (
	"gorm.io/gorm"
	"pheasant-api/database"
	"time"
)

type UserModel interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
}

type userModelDependencies struct {
	db *gorm.DB
}

// User is the main user model.
type User struct {
	ID        uint64    `json:"-" gorm:"primaryKey"`
	UUID      string    `json:"uuid" gorm:"index:,unique; default: uuid_generate_v4()"`
	Email     string    `json:"email" gorm:"index:,unique"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUserModel() UserModel {
	return &userModelDependencies{
		db: database.DB,
	}
}

func (model *userModelDependencies) GetUserByEmail(email string) (*User, error) {
	user := User{}
	userResult := model.db.Where("email = ?", email).First(&user)
	if userResult.Error != nil {
		return nil, userResult.Error
	}

	return &user, nil
}

func (model *userModelDependencies) CreateUser(user *User) (*User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	userResult := model.db.Create(user)
	if userResult.Error != nil {
		return nil, userResult.Error
	}

	return user, nil
}
