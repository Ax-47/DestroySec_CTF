package wawha

import (
	"api/jwt/service"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context, tag, username, OBJ string) string {
	JWt, _ := service.GenerateTokenNEW(c, tag, username, OBJ)
	return JWt
}
