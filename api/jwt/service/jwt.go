package service

import (
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//jwt service
func GenerateToken(c *gin.Context, key string, stats string, otp int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		&jwt.StandardClaims{
			Audience:  key,
			IssuedAt:  otp,
			Subject:   stats,
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		})

	ss, err := token.SignedString([]byte("MySignature"))

	return ss, err
}
func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("MySignature"), nil
	})

	return err
}
func DecodeToken(token string) (*jwt.Token, error) {
	kk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("MySignature"), nil
	})
	return kk, err
}
