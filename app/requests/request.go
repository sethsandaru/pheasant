package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func validate(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(&body); err != nil {
		validateObj := validator.New()

		// validation error
		if err := validateObj.Struct(body); err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{
					"error":    "Validation Error",
					"messages": err.Error(),
				},
			)

			return err
		}

		// binding error
		return err
	}

	return nil
}
