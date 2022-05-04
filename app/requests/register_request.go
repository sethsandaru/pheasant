package requests

import "github.com/gin-gonic/gin"

type RegisterRequest interface {
	Validate(c *gin.Context) RegisterBody
}

type registerRequestDependencies struct{}

type RegisterBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

func GetRegisterRequest() RegisterRequest {
	return &registerRequestDependencies{}
}

func (request *registerRequestDependencies) Validate(c *gin.Context) RegisterBody {
	registerBody := RegisterBody{}

	validate(c, &registerBody)

	return registerBody
}
