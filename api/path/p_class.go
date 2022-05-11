package path

import (
	db "api/db"
	h "api/hashpaww"
	"time"

	"math/rand"

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

type ln struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Reg struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}
type DATA struct {
	Email                 string
	Username, Tag, UserId string
	Subdata               struct {
		Password string
	}
}

func GanuserTag(s db.Db_mongo) string {
	for {
		bytes := make([]byte, 4)
		var pool = "1234567890"
		for i := 0; i < 4; i++ {
			bytes[i] = pool[rand.Intn(len(pool))]
		}
		some, _ := s.Db_FindALL("userid", string(bytes))
		if some == nil {
			return string(bytes)
		}
	}

}
func Ganuserid(s db.Db_mongo) string {
	for {
		var pool = "1234567890"
		dd := make([]byte, 13)
		for i := 0; i < 13; i++ {
			dd[i] = pool[rand.Intn(len(pool))]
		}
		some, _ := s.Db_FindALL("userid", string(dd))
		if some == nil {
			return string(dd)
		}
	}

}
func Register(c *gin.Context) {
	var fromreg Reg
	if err := c.BindJSON(&fromreg); err != nil {
		return
	}
	password := fromreg.Password
	repassword := fromreg.Repassword
	email := fromreg.Email
	user := fromreg.Username
	if user != "" || email != "" || repassword != "" || password != "" {
		if password == repassword {
			var s db.Db_mongo
			s.Db_start()
			ch := s.Db_FindtOne("email", email)

			if ch == nil {

				var post DATA
				post.Email = email
				post.Subdata.Password = h.Mhash(password)
				post.Username = user
				post.UserId = string(Ganuserid(s))
				post.Tag = string(GanuserTag(s))
				s.Db_InsertOneS(post)
				const userkey = "email"
				const timee = "time"
				session := sessions.Default(c)
				cookie := email
				time := time.Now()
				session.Set(userkey, time)
				session.Set(timee, cookie)
				if err := session.Save(); err != nil {
					c.JSON(200, gin.H{"error": "Failed to save session"})
					return
				}
				c.JSON(200, gin.H{
					"message": "Register suss",
				})
			} else {
				c.JSON(504, gin.H{
					"message": "have email",
				})
			}

		} else {
			c.JSON(404, gin.H{
				"message": "password!=repassword fail",
			})
		}
	} else {
		c.JSON(404, gin.H{
			"message": "em fail",
		})
	}
}
func Login(c *gin.Context) {
	session := sessions.Default(c)
	var fromreg ln
	if err := c.BindJSON(&fromreg); err != nil {
		return
	}
	email := fromreg.Email
	password := fromreg.Password

	if email != "" || password != "" {
		var s db.Db_mongo
		s.Db_start()
		key := s.Db_FindtOne("username", email)
		if key != nil {
			kpass := key[2]
			has := kpass.Value.(string)
			const userkey = "email"
			const timee = "time"
			if h.Vcheck(has, password) {

				if err := session.Save(); err != nil {
					c.JSON(200, gin.H{"error": "Failed to save session"})
					return
				}
				cookie := email
				time := time.Now()
				session.Set(userkey, time)
				session.Set(timee, cookie)
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
		} else {

			c.JSON(404, gin.H{
				"message": "email not font"})
		}
	} else {

		c.JSON(404, gin.H{
			"message": "em fail"})
	}
}
