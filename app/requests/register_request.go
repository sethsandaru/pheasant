package requests

import "github.com/gin-gonic/gin"

type RegisterRequest interface {
	Validate(c *gin.Context) (RegisterBody, error)
}

type registerRequestDependencies struct{}

type RegisterBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}

func GetRegisterRequest() RegisterRequest {
	return &registerRequestDependencies{}
}

func (request *registerRequestDependencies) Validate(c *gin.Context) (RegisterBody, error) {
	registerBody := RegisterBody{}

	err := validate(c, &registerBody)

	return registerBody, err
}
