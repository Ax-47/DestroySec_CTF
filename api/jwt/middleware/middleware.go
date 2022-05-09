package middleware

import (
	"api/jwt/service"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		tokenString := authHeader[len(BEARER_SCHEMA):]

		if token, err := service.JWTAuthService().ValidateToken(tokenString); err != nil {
			c.JSON(200, gin.H{"error": "Failed to save session"})
			return
		} else {
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println(claims)
			} else {
				fmt.Println("testing")
				fmt.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}

	}
}
