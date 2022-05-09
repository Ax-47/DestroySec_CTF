package main

import (
	p "api/path"
	"errors"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	limiter "github.com/julianshen/gin-limiter"
)

func main() {
	var secret = []byte("secret")

	r := gin.Default()
	lm := limiter.NewRateLimiter(time.Minute, 10, func(ctx *gin.Context) (string, error) {
		key := ctx.Request.Header.Get("X-API-KEY")
		if key != "" {
			return key, nil
		}
		return "", errors.New("API key is missing")
	})
	r.Use(sessions.Sessions("DestroySce", sessions.NewCookieStore(secret)))
	r.POST("/q", p.M)
	r.POST("/register", lm.Middleware(), p.Register)
	r.POST("/login", lm.Middleware(), p.Login)

	v1 := r.Group("/get")

	{
		v1.GET("/api", func(ctx *gin.Context) {
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
