package path

import (
	db "api/db"
	p "api/path/login"
	P "api/path/register"

	//"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var s db.Db_mongo

func init() {
	s.Db_start()
}

func M(c *gin.Context) {

	Leve := c.PostForm("leve")
	Asm := c.PostForm("asm")

	if Leve == "1" {
		if Asm == "hackerman" {
			c.JSON(200, gin.H{
				"message": Asm + "leve:" + Leve + " True",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "leve:" + Leve + " False",
			})
		}
	}

}
func C(c *gin.Context) {
	session := sessions.Default(c)

	v := session.Get("user")

	g, _ := 11, 22

	session.Set("user", g)
	session.Save()
	c.JSON(200, gin.H{"count": v})
}

func Register(c *gin.Context) {

	P.Register(c, s)
}
func Login(c *gin.Context) {

	p.Login(c, s)
}
