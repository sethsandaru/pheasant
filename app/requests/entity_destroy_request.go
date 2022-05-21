package requests

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
)

type EntityDestroyRequest interface {
	Authorize(c *gin.Context) bool
	GetEntity(c *gin.Context) *models.Entity
}

type entityDestroyRequestDependencies struct{}

func GetEntityDestroyRequest() EntityDestroyRequest {
	return &entityDestroyRequestDependencies{}
}

func (request *entityDestroyRequestDependencies) Authorize(c *gin.Context) bool {
	return GetEntityShowRequest().Authorize(c)
}

func (request *entityDestroyRequestDependencies) GetEntity(c *gin.Context) *models.Entity {
	return GetEntityShowRequest().GetEntity(c)
}
