package controllers

import (
	"github.com/gin-gonic/gin"
)

type ReleaseController interface {
	Index(c *gin.Context)
}

type releaseControllerParams struct{}

func GetReleaseController() ReleaseController {
	return &releaseControllerParams{}
}

// Index Get a list of Releases
func (controller *releaseControllerParams) Index(c *gin.Context) {
	respondOk(c, "OK")
}
