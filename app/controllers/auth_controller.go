package controllers

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/requests"
	"pheasant-api/app/services"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	ForgotPassword(c *gin.Context)
}

type authControllerDependencies struct {
	authService services.AuthService
}

func GetAuthController() AuthController {
	return &authControllerDependencies{
		authService: services.GetAuthService(),
	}
}

func (controller *authControllerDependencies) Login(c *gin.Context) {
	loginBody := requests.ValidateLoginRequest(c)
	token, err := controller.authService.CheckAuthentication(loginBody)
	if err != nil {
		respondBadRequest(c, gin.H{
			"error": "Email or Password is wrong, please check your detail again.",
		})
	}

	// ensure token created successfully
	if token == "" {
		respondInternalServerError(c, gin.H{
			"error": "Failed to login due to internal server error, please try again.",
		})
	}

	respondOk(c, gin.H{
		"token": token,
	})
}

func (controller *authControllerDependencies) Register(c *gin.Context) {

}

func (controller *authControllerDependencies) ForgotPassword(c *gin.Context) {

}
