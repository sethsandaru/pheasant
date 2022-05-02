package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func respondOk(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func respondCreated(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusCreated, obj)
}

func respondBadRequest(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusBadRequest, obj)
}

func respondUnprocessableEntity(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusUnprocessableEntity, obj)
}

func respondInternalServerError(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusInternalServerError, obj)
}

func respondServiceUnavailable(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusServiceUnavailable, obj)
}
