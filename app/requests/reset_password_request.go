package requests

import "github.com/gin-gonic/gin"

type ResetPasswordRequest interface {
	Validate(c *gin.Context) (ResetPasswordBody, error)
}

type resetPasswordRequestDependencies struct{}

type ResetPasswordBody struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func GetResetPasswordRequest() ResetPasswordRequest {
	return &resetPasswordRequestDependencies{}
}

func (request *resetPasswordRequestDependencies) Validate(c *gin.Context) (ResetPasswordBody, error) {
	resetPasswordBody := ResetPasswordBody{}

	err := validate(c, &resetPasswordBody)

	return resetPasswordBody, err
}
