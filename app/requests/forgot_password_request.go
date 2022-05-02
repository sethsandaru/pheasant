package requests

import "github.com/gin-gonic/gin"

type ForgotPasswordRequest interface {
	Validate(c *gin.Context) ForgotPasswordBody
}

type forgotPasswordRequestDependencies struct{}

type ForgotPasswordBody struct {
	Email string `json:"email" binding:"required,email"`
}

func GetForgotPasswordRequest() ForgotPasswordRequest {
	return &forgotPasswordRequestDependencies{}
}

func (request *forgotPasswordRequestDependencies) Validate(c *gin.Context) ForgotPasswordBody {
	forgotPasswordBody := ForgotPasswordBody{}

	validate(c, &forgotPasswordBody)

	return forgotPasswordBody
}
