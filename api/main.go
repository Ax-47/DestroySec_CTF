package main

import (
	"api/middleware"
	p "api/path"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	var secret = []byte("secret")
	r := gin.Default()
	r.Use(sessions.Sessions("DestroySce", sessions.NewCookieStore(secret)))
	r.POST("/q", p.M)
	r.POST("/register", p.Register)
	r.POST("/login", p.Login)
	v1 := r.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})
	}
	r.Run(":8080")

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	/*
		#Example
		import requets
		requets.post("localhost:8080/q",data={"leve":"1","ams":"hackerman"})

	*/

}
