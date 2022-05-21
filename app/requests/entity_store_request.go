package requests

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
)

type EntityStoreRequest interface {
	Validate(c *gin.Context) (EntityStoreBody, error)
	GetUser(c *gin.Context) *models.User
}

type entityStoreRequestDependencies struct{}

type EntityStoreBody struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func GetEntityStoreRequest() EntityStoreRequest {
	return &entityStoreRequestDependencies{}
}

func (request *entityStoreRequestDependencies) Validate(c *gin.Context) (EntityStoreBody, error) {
	entityStoryBody := EntityStoreBody{}

	err := validate(c, &entityStoryBody)

	return entityStoryBody, err
}

func (request *entityStoreRequestDependencies) GetUser(c *gin.Context) *models.User {
	return getUser(c)
}
