package controllers

import "github.com/gin-gonic/gin"

func respondOk(c *gin.Context, obj interface{}) {
	c.JSON(200, obj)
}

func respondCreated(c *gin.Context, obj interface{}) {
	c.JSON(201, obj)
}

func respondBadRequest(c *gin.Context, obj interface{}) {
	c.JSON(400, obj)
}

func respondUnprocessedEntity(c *gin.Context, obj interface{}) {
	c.JSON(422, obj)
}

func respondInternalServerError(c *gin.Context, obj interface{}) {
	c.JSON(500, obj)
}

func respondServiceUnavailable(c *gin.Context, obj interface{}) {
	c.JSON(503, obj)
}
