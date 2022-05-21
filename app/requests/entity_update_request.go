package requests

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
)

type EntityUpdateRequest interface {
	Authorize(c *gin.Context) bool
	GetEntity(c *gin.Context) *models.Entity
	Validate(c *gin.Context) (EntityUpdateBody, error)
}

type entityUpdateRequestDependencies struct{}

func GetEntityUpdateRequest() EntityUpdateRequest {
	return &entityUpdateRequestDependencies{}
}

type EntityUpdateBody struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (request *entityUpdateRequestDependencies) Authorize(c *gin.Context) bool {
	return GetEntityShowRequest().Authorize(c)
}

func (request *entityUpdateRequestDependencies) GetEntity(c *gin.Context) *models.Entity {
	return GetEntityShowRequest().GetEntity(c)
}

func (request *entityUpdateRequestDependencies) Validate(c *gin.Context) (EntityUpdateBody, error) {
	entityUpdateBody := EntityUpdateBody{}

	err := validate(c, &entityUpdateBody)

	return entityUpdateBody, err
}
