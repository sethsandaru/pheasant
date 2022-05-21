package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pheasant-api/app/models"
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
		tokenCheck, err, claims := services.GetAuthService().ValidateToken(token)
		if err != nil || !tokenCheck.Valid {
			abortUnauthenticated(c)

			return
		}

		// inject user instance across the request lifecycle
		user, err := models.GetUserModel().Find(claims.UserId)
		if err != nil {
			abortUnauthenticated(c)

			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func abortUnauthenticated(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthenticated",
	})
}
