package path

import (
	db "api/db"
	h "api/hashpaww"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

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
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	if password == repassword {
		var masp = make(map[string]string)
		masp["username"] = username
		masp["password"] = h.Mhash(password)
		var s db.Db_mongo
		s.Db_start()
		s.Db_InsertOne(masp)
		const userkey = "user"
		session := sessions.Default(c)
		session.Set(userkey, username)
		if err := session.Save(); err != nil {
			c.JSON(200, gin.H{"error": "Failed to save session"})
			return
		}
		c.JSON(200, gin.H{
			"message": "Register suss",
		})

	} else {
		c.JSON(404, gin.H{
			"message": "Register fail",
		})
	}
}
func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	var s db.Db_mongo
	s.Db_start()

	key := s.Db_FindtOne(username)
	kpass := key[2]
	has := kpass.Value.(string)
	const userkey = "user"
	if h.Vcheck(has, password) {
		session.Set(userkey, username)
		if err := session.Save(); err != nil {
			c.JSON(200, gin.H{"error": "Failed to save session"})
			return
		}
		c.JSON(200, gin.H{
			"message": "login suss",
		})
	} else {

		c.JSON(404, gin.H{
			"message": "login fail"})
	}
}
