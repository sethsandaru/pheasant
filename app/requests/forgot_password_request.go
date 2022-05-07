package requests

import "github.com/gin-gonic/gin"

type ForgotPasswordRequest interface {
	Validate(c *gin.Context) (ForgotPasswordBody, error)
}

type forgotPasswordRequestDependencies struct{}

type ForgotPasswordBody struct {
	Email string `json:"email" validate:"required,email"`
}

func GetForgotPasswordRequest() ForgotPasswordRequest {
	return &forgotPasswordRequestDependencies{}
}

func (request *forgotPasswordRequestDependencies) Validate(c *gin.Context) (ForgotPasswordBody, error) {
	forgotPasswordBody := ForgotPasswordBody{}

	err := validate(c, &forgotPasswordBody)

	return forgotPasswordBody, err
}
