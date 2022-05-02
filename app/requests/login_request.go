package requests

import "github.com/gin-gonic/gin"

type LoginBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func ValidateLoginRequest(c *gin.Context) LoginBody {
	loginBody := LoginBody{}

	validate(c, &loginBody)

	return loginBody
}
