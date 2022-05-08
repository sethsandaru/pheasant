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
	FindByUuid(uuid string) (*Entity, error)
	Create(entity Entity) (*Entity, error)
	Update(entity Entity) (*Entity, error)
	Delete(entity Entity) bool
}

type entityModelDependencies struct{}

func GetEntityModel() EntityModel {
	return &entityModelDependencies{}
}

func (model *entityModelDependencies) FindByUuid(uuid string) (*Entity, error) {
	entity := &Entity{}
	findResult := findByUuidQuery(uuid).First(entity)
	if findResult.Error != nil {
		return nil, findResult.Error
	}

	return entity, nil
}

func (model *entityModelDependencies) Create(entity Entity) (*Entity, error) {
	createResult := DB.Create(&entity)
	if createResult.Error != nil {
		return nil, createResult.Error
	}

	return &entity, nil
}

func (model *entityModelDependencies) Update(entity Entity) (*Entity, error) {
	saveResult := DB.Save(&entity)
	if saveResult.Error != nil {
		return nil, saveResult.Error
	}

	return &entity, nil
}

func (model *entityModelDependencies) Delete(entity Entity) bool {
	deleteResult := DB.Delete(&entity)
	return deleteResult.Error == nil
}
