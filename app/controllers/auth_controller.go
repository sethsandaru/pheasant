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
	ValidateResetPasswordToken(c *gin.Context)
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
	loginBody, err := requests.GetLoginRequest().Validate(c)
	if err != nil {
		return
	}

	token, err := controller.authService.CheckAuthentication(loginBody)
	if err != nil {
		respondBadRequest(c, gin.H{
			"error": "Email or Password is wrong, please check your detail again.",
		})

		return
	}

	// ensure token created successfully
	if token == "" {
		respondInternalServerError(c, gin.H{
			"error": "Failed to login due to internal server error, please try again.",
		})

		return
	}

	respondOk(c, gin.H{
		"token": token,
	})
}

func (controller *authControllerDependencies) Register(c *gin.Context) {
	registerBody, err := requests.GetRegisterRequest().Validate(c)
	if err != nil { // can't abort from the request, so this is the only way T.T
		return
	}

	// register
	user, err := controller.authService.Register(
		registerBody.Email,
		registerBody.Password,
		registerBody.FullName,
	)
	if err != nil {
		respondBadRequest(c, gin.H{
			"error": err.Error(),
		})

		return
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

		return
	}

	respondOk(c, gin.H{
		"status": true,
	})
}

func (controller *authControllerDependencies) ValidateResetPasswordToken(c *gin.Context) {
	token := c.Param("token")
	isTokenValid := controller.authService.IsResetPasswordTokenStillValid(token)

	if isTokenValid {
		respondOk(c, gin.H{
			"valid": true,
		})
	} else {
		respondBadRequest(c, gin.H{
			"valid": false,
		})
	}
}

func (controller *authControllerDependencies) ResetPassword(c *gin.Context) {
	token := c.Param("token")
	isTokenValid := controller.authService.IsResetPasswordTokenStillValid(token)
	if !isTokenValid {
		respondBadRequest(c, gin.H{
			"message": "Invalid reset password token",
		})
	}

	resetPasswordBody, err := requests.GetResetPasswordRequest().Validate(c)
	if err != nil {
		return
	}

	// reset password here
	err = controller.authService.ResetPassword(token, resetPasswordBody.Password)
	if err != nil {
		respondBadRequest(c, gin.H{
			"message": err.Error(),
		})
	} else {
		respondOk(c, gin.H{
			"message": "Reset password successfully. You can login again now.",
		})
	}
}
