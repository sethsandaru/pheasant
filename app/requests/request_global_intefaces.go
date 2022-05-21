package requests

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
)

type GetUserInterface interface {
	GetUser(c *gin.Context) *models.User
}

type HasAuthorizationInterface interface {
	Authorize(c *gin.Context) bool
}
