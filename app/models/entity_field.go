package models

import (
	"database/sql"
	"time"
)

const (
	FieldTypeString   string = "string"
	FieldTypeInteger         = "integer"
	FieldTypeDouble          = "double"
	FieldTypeText            = "text"
	FieldTypeDate            = "date"
	FieldTypeTime            = "time"
	FieldTypeDateTime        = "datetime"
)

type EntityField struct {
	ID uint64 `json:"-" gorm:"primaryKey"`
	HasUUID
	EntityID    uint64         `json:"entity_id" gorm:"index"`
	Title       string         `json:"title" gorm:"index"`
	Description sql.NullString `json:"description"`
	Type        string         `json:"type"`
	DefaultData sql.NullString `json:"default_data"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"index"`

	Entity Entity `json:"entity" gorm:"foreignKey:EntityID"`
}

type EntityFieldModel interface {
	FindByUuid(uuid string) (*EntityField, error)
}

type entityFieldDependencies struct{}

func GetEntityFieldModel() EntityFieldModel {
	return &entityFieldDependencies{}
}

func (model *entityFieldDependencies) FindByUuid(uuid string) (*EntityField, error) {
	entity := &EntityField{}
	findResult := findByUuidQuery(uuid).First(entity)
	if findResult.Error != nil {
		return nil, findResult.Error
	}

	return entity, nil
}
