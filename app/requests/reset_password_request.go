package requests

import "github.com/gin-gonic/gin"

type ResetPasswordRequest interface {
	Validate(c *gin.Context) (ResetPasswordBody, error)
}

type resetPasswordRequestDependencies struct{}

type ResetPasswordBody struct {
	Password string `json:"password" binding:"required"`
}

func GetResetPasswordRequest() ResetPasswordRequest {
	return &resetPasswordRequestDependencies{}
}

func (request *resetPasswordRequestDependencies) Validate(c *gin.Context) (ResetPasswordBody, error) {
	resetPasswordBody := ResetPasswordBody{}

	err := validate(c, &resetPasswordBody)

	return resetPasswordBody, err
}
