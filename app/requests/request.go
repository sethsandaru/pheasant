package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pheasant-api/app/models"
)

func validate(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"error": "Binding Error. Make sure your content is JSON",
			},
		)

		// binding error
		return err
	}

	validateObj := validator.New()
	if validateErr := validateObj.Struct(body); validateErr != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"error":    "Validation Error",
				"messages": validateErr.Error(),
			},
		)

		return validateErr
	}

	return nil
}

func getUser(c *gin.Context) *models.User {
	userRequest, isExists := c.Get("user")
	if !isExists {
		return nil
	}

	user := userRequest.(*models.User)
	return user
}

func abortUnauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(
		http.StatusForbidden,
		gin.H{"error": "This action is unauthorized."},
	)
}
