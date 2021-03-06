package main

import (
	//jwt "api/jwt/service"
	M "api/Middleware"
	p "api/path/compile_path"
	"errors"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	limiter "github.com/julianshen/gin-limiter"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,X-API-KEY, jwt")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("destroy", store))
	//sessionNames := []string{"DestorySce", "Just kidding"}

	r.Use(CORSMiddleware())
	lm := limiter.NewRateLimiter(time.Minute, 10, func(ctx *gin.Context) (string, error) {
		key := ctx.Request.Header.Get("X-API-KEY")
		if key != "" {
			return key, nil
		}
		return "", errors.New("API key is missing")
	})
	api := r.Group("/", M.AuthorizationMiddleware)
	apilogin := r.Group("/apilogin")
	apilogin.POST("/ln", lm.Middleware(), p.Login) //fix this
	apilogin.POST("/reg", lm.Middleware(), p.Register)
	api.POST("/q", p.M)

	apilogin.GET("/Check", p.C)
	r.Run(":9000")

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	/*
		#Example
		import requets
		requets.post("localhost:8080/q",data={"leve":"1","ams":"hackerman"})

	*/

}
