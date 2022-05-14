package requests

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
)

type EntityShowRequest interface {
	Authorize(c *gin.Context) bool
	GetEntity(c *gin.Context) *models.Entity
}

type entityShowRequestDependencies struct{}

func GetEntityShowRequest() EntityShowRequest {
	return &entityShowRequestDependencies{}
}

func (request *entityShowRequestDependencies) Authorize(c *gin.Context) bool {
	user := getUser(c)
	entity := request.GetEntity(c)

	if user.ID != entity.UserID {
		abortUnauthorized(c)

		return false
	}

	return true
}

func (request *entityShowRequestDependencies) GetEntity(c *gin.Context) *models.Entity {
	entityRequest, _ := c.Get("entity")
	entity := entityRequest.(*models.Entity)

	return entity
}
