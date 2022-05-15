package path

import (
	db "api/db"

	p "api/path/login"
	P "api/path/register"
	"math/rand"

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

	P.Register(c, s)
}
func Login(c *gin.Context) {

	p.Login(c, s)
}
