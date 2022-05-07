package requests

import "github.com/gin-gonic/gin"

type LoginRequest interface {
	Validate(c *gin.Context) (LoginBody, error)
}

type loginRequestDependencies struct{}

type LoginBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func GetLoginRequest() LoginRequest {
	return &loginRequestDependencies{}
}

func (request *loginRequestDependencies) Validate(c *gin.Context) (LoginBody, error) {
	loginBody := LoginBody{}

	err := validate(c, &loginBody)

	return loginBody, err
}
