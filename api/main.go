package main

import (
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
	r.Run(":8080")

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	/*
		#Example
		import requets
		requets.post("localhost:8080/q",data={"leve":"1","ams":"hackerman"})

	*/

}
