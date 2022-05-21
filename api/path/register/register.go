package register

import (
	db "api/db"
	h "api/hash_class"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func saveDAta(s db.Db_mongo, email, password, user string) {
	var post DATA
	post.Email = email
	post.Subdata.Password = h.Mhash(password)
	post.Username = user
	post.UserId = string(Ganuserid(s))
	post.Tag = string(GanuserTag(s))
	go s.Db_InsertOneS(post)
}

func Register(c *gin.Context, s db.Db_mongo) {
	var fromreg Reg
	if err := c.BindJSON(&fromreg); err != nil {
		return
	}
	password := fromreg.Password
	repassword := fromreg.Repassword
	email := fromreg.Email
	user := fromreg.Username
	cha := make(chan primitive.D)
	go s.Db_FindtOne("email", email, cha)
	if user != "" || email != "" || repassword != "" || password != "" {
		if password == repassword {
			datauser := <-cha
			if datauser == nil {
				go saveDAta(s, email, password, user)
				c.JSON(200, gin.H{
					"message": "Register suss",
				})
			} else {
				c.JSON(404, gin.H{
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
