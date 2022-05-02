package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func validate(c *gin.Context, body interface{}) {
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"error":    "Validation Error",
				"messages": err.Error(),
			},
		)
	}
}
