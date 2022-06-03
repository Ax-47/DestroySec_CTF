package Middleware

import (
	jwt "api/jwt/service"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	var request DataJwT
	err := c.ShouldBindHeader(&request)
	if err != nil {
		fmt.Print(err)
		return
	}
	s := c.Request.Header.Get("jwt")

	token := strings.TrimPrefix(s, "Bearer ")
	//_ , err := jwt.DecodeToken(token)
	fmt.Println(request)
	fmt.Println(s)
	if err := jwt.ValidateToken(token); err != nil {
		c.AbortWithStatus(505)
		return
	}
}
