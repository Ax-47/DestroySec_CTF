package khawmankai

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func s(c *gin.Context) {
	// Get cookie key_ value corresponding to cookie
	cookie_value, err := c.Cookie("my_cookie")

	if err != nil {
		cookie_value = "DestroySec"
		// cookie settings
		c.SetCookie("my_cookie", cookie_value, 3600, "/", "localhost", false, true)
	}

	// Then we can look in the response header. result is the key of the response header and the cookie_value is the value of the response header
	c.Header("result", cookie_value)
}
func Check(c *gin.Context) {
	// Get cookie key_ value corresponding to cookie

	s := sessions.Default(c)
	s.Save()
	sessionID := s.Get("jwt")
	fmt.Print(sessionID.(string))
	c.JSON(200, gin.H{
		"message": "login suss",
		"jwt":     sessionID.(string),
	})
	// Then we can look in the response header. result is the key of the response header and the cookie_value is the value of the response header
	//c.Header("result", )
}
