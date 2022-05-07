package controllers

import (
	"github.com/gin-gonic/gin"
)

type EntityController interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type entityControllerDependencies struct{}

func GetEntityController() EntityController {
	return &entityControllerDependencies{}
}

// Index Get a list of Entities of User
func (controller *entityControllerDependencies) Index(c *gin.Context) {
	respondOk(c, "OK")
}

func (controller *entityControllerDependencies) Show(c *gin.Context) {
	respondOk(c, "OK")
}

func (controller *entityControllerDependencies) Store(c *gin.Context) {
	respondOk(c, "OK")
}

func (controller *entityControllerDependencies) Update(c *gin.Context) {
	respondOk(c, "OK")
}

func (controller *entityControllerDependencies) Destroy(c *gin.Context) {
	respondOk(c, "OK")
}
