package routes

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/controllers"
	"pheasant-api/app/middlewares"
)

// SetupApiRouter sets up the router.
func SetupApiRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		registerAuthenticationRoutes(v1)

		v1Auth := v1.Group("/")
		{
			v1Auth.Use(middlewares.RequiresAuth())
			registerReleaseRoutes(v1Auth)
			registerEntityRoutes(v1Auth)
		}
	}

	return router
}

// auth-group routes
func registerAuthenticationRoutes(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")
	{
		authController := controllers.GetAuthController()

		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.POST("/forgot-password", authController.ForgotPassword)
		auth.GET("/reset-password-token/:token", authController.ValidateResetPasswordToken)
		auth.POST("/reset-password", authController.ResetPassword)
	}
}

// release-group routes
func registerReleaseRoutes(v1 *gin.RouterGroup) {
	release := v1.Group("/releases")
	{
		releaseController := controllers.GetReleaseController()

		release.GET("", releaseController.Index)
	}
}

// entity-group routes
func registerEntityRoutes(v1 *gin.RouterGroup) {
	entity := v1.Group("/entities")
	{
		entityController := controllers.GetEntityController()

		entity.GET("", entityController.Index)
		entity.GET("/:entityUuid", entityController.Show)
		entity.POST("", entityController.Store)
		entity.PUT("/:entityUuid", entityController.Update)
		entity.DELETE("/:entityUuid", entityController.Destroy)
	}
}
