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
	ResetPassword(c *gin.Context)
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
	loginBody := requests.GetLoginRequest().Validate(c)
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
	registerBody := requests.GetRegisterRequest().Validate(c)
	user, err := controller.authService.Register(
		registerBody.Email,
		registerBody.Password,
		registerBody.FullName,
	)
	if err != nil {
		respondBadRequest(c, gin.H{
			"error": "Register failed, please try again",
		})
	}

	respondCreated(c, user)
}

func (controller *authControllerDependencies) ForgotPassword(c *gin.Context) {
	forgotPasswordBody := requests.GetForgotPasswordRequest().Validate(c)
	processStatus := controller.authService.ForgotPassword(forgotPasswordBody.Email)
	if !processStatus {
		respondBadRequest(c, gin.H{
			"error": "Email didn't exists on our system or internal error happened. Please check and try again.",
		})
	}

	respondOk(c, gin.H{
		"status": true,
	})
}

func (controller *authControllerDependencies) ResetPassword(c *gin.Context) {
	// WIP
}
