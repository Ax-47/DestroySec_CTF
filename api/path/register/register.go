package register

import (
	db "api/db"
	"api/gmail"
	h "api/hash_class"
	jwt "api/jwt/service"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reg struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}
type DATA struct {
	Email                       string
	Username, Tag, UserId, Time string
	SessionReg                  string
	Identifind                  bool
	Subdata                     struct {
		Password string
	}
}

func saveDAta(s db.Db_mongo, email, password, user, tag, SEss, time string) {
	var post DATA
	post.Time = time
	post.Email = email
	post.Subdata.Password = h.Mhash(password)
	post.Username = user
	post.UserId = string(Ganuserid(s))
	post.SessionReg = SEss
	post.Tag = tag
	post.Identifind = false
	go s.Db_InsertOneS_UniDentify(post)
}

func Register(c *gin.Context, s db.Db_mongo, am gmail.GAmll) {
	var fromreg Reg
	if err := c.BindJSON(&fromreg); err != nil {
		return
	}
	password := fromreg.Password
	repassword := fromreg.Repassword
	email := fromreg.Email
	user := fromreg.Username
	cha := make(chan primitive.D)
	sha := make(chan primitive.D)
	go s.Db_FindtOne("email", email, cha)
	go s.Db_FindtOne_UniDentify("email", email, sha)
	if user != "" || email != "" || repassword != "" || password != "" {
		if password == repassword {
			datauser := <-cha
			reguserunidentify := sha

			if datauser == nil || datauser.Map()["Identifind"] == false {
				if reguserunidentify != nil {

					s.Db_DeleteMany_UniDentify(bson.M{"email": bson.M{"$eq": email}})

					tag := string(GanuserTag(s))
					Ax := GenOTP()
					t := time.Now().Format("2006-01-02 15:04:05")
					g, _ := jwt.GenerateTokenReg(c, tag, user, email, t, int64(60456))
					hg := h.Mhash(g)
					hAx := h.Mhash(Ax)
					go saveDAta(s, email, password, user, tag, hg+" "+hAx, t)

					am.SEndlogin(user, tag, Ax, email)
					c.JSON(200, gin.H{
						"message": "Register suss",
						"hee":     g,
					})
				} else {
					c.JSON(403, gin.H{
						"message": "have email",
					})
				}

			} else {
				c.JSON(402, gin.H{
					"message": "password!=repassword fail",
				})
			}
		} else {
			c.JSON(401, gin.H{
				"message": "em fail",
			})
		}
	}
}
