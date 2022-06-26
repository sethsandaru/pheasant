package controllers

import (
	"github.com/gin-gonic/gin"
)

type EntityFieldController interface {
	Index(c *gin.Context)
	Update(c *gin.Context)
}

type entityFieldControllerParams struct{}

func GetEntityFieldController() EntityFieldController {
	return &entityFieldControllerParams{}
}

// Index Get a list of EntityFields
func (controller *entityFieldControllerParams) Index(c *gin.Context) {
	respondOk(c, "OK")
}

// Update will update a list of EntityFields
func (controller *entityFieldControllerParams) Update(c *gin.Context) {
	respondOk(c, "OK")
}
