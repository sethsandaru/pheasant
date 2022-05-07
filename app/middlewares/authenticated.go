package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pheasant-api/app/services"
)

func RequiresAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			abortUnauthenticated(c)

			return
		}

		token := authHeader[len("Bearer "):]
		tokenCheck, err := services.GetAuthService().ValidateToken(token)
		if err != nil || !tokenCheck.Valid {
			abortUnauthenticated(c)

			return
		}

		c.Next()
	}
}

func abortUnauthenticated(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthenticated",
	})
}
