package controllers

import (
	"github.com/gin-gonic/gin"
)

// IndexRelease Get a list of Releases
func IndexRelease(c *gin.Context) {
	respondOk(c, "OK")
}
