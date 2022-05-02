package routes

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/controllers"
)

// SetupApiRouter sets up the router.
func SetupApiRouter() *gin.Engine {
	r := gin.Default()

	r.Group("/auth")
	{
		r.POST("/login", controllers.Login)
		r.POST("/register", controllers.Register)
		r.POST("/forgot-password", controllers.ForgotPassword)
	}

	r.Group("/releases")
	{
		r.POST("/", controllers.IndexRelease)
	}

	return r
}
