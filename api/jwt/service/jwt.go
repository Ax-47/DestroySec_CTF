package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//jwt service
func GenerateToken(c *gin.Context, key string, otp int64) (string,error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		&jwt.StandardClaims{
			Audience:  key,
			IssuedAt:  otp,
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		})

	ss, err := token.SignedString([]byte("MySignature"))
	return ss, err
}
func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("MySignature"), nil
	})

	return err
}
func AuthorizeJWT(c *gin.Context) {

	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
