package path

import (
	db "api/db"
	"api/gmail"
	p "api/path/login"
	P "api/path/register"
	v "api/path/verify_gmail"
	vu "api/path/verifyuser"

	//"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var s db.Db_mongo
var am gmail.GAmll

func init() {
	s.Db_start()
	am.Login("axc47chaos@gmail.com", "eaighfojzsjfhtda")
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
func Verifyotp_func(c *gin.Context) {
	v.Verifyotp(c, s)

}
func Verifyotp_Reg_func(c *gin.Context) {
	vu.Verifyotp(c, s)

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

	P.Register(c, s, am)
}
func Login(c *gin.Context) {

	p.Login(c, s, am)
}
