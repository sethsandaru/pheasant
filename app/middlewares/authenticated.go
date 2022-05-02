package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pheasant-api/app/services"
)

func RequiresAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authentication")
		token := authHeader[len("Bearer"):]

		tokenCheck, err := services.GetAuthService().ValidateToken(token)
		if err != nil || !tokenCheck.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
