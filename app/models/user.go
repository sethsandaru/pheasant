package models

import (
	"errors"
	"log"
	"time"
)

type UserModel interface {
	GetUserByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
}

type userModelDependencies struct{}

// User is the main user model.
type User struct {
	ID        uint64    `json:"-" gorm:"primaryKey"`
	UUID      string    `json:"uuid" gorm:"index:,unique; default: uuid_generate_v4()"`
	Email     string    `json:"email" gorm:"index:,unique"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"index"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}

func GetUserModel() UserModel {
	return &userModelDependencies{}
}

func (model *userModelDependencies) GetUserByEmail(email string) (*User, error) {
	user := User{}
	userResult := DB.Where("email = ?", email).First(&user)
	if userResult.Error != nil {
		return nil, userResult.Error
	}

	return &user, nil
}

func (model *userModelDependencies) Create(user *User) (*User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	userResult := DB.Create(user)
	if userResult.Error != nil {
		log.Print(userResult.Error)
		return nil, errors.New("Failed to create new user")
	}

	return user, nil
}

func (model *userModelDependencies) Update(user *User) (*User, error) {
	user.UpdatedAt = time.Now()

	userResult := DB.Save(user)
	if userResult.Error != nil {
		log.Print(userResult.Error)
		return nil, errors.New("Failed to update user")
	}

	return user, nil
}
