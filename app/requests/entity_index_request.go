package requests

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
)

type EntityIndexRequest interface {
	GetUserInterface
}

type entityIndexRequestDependencies struct{}

func GetEntityIndexRequest() EntityIndexRequest {
	return &entityIndexRequestDependencies{}
}

func (request *entityIndexRequestDependencies) GetUser(c *gin.Context) *models.User {
	return getUser(c)
}
