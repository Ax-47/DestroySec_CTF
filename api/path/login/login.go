package login

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

type DATA struct {
	Sessions_DATA struct {
		session string
	}
}

func Login(c *gin.Context, s db.Db_mongo, am gmail.GAmll) {

	var fromreg ln
	if err := c.BindJSON(&fromreg); err != nil {
		return
	}
	email := fromreg.Email
	cha := make(chan primitive.D)
	go s.Db_FindtOne("email", email, cha)
	password := fromreg.Password

	if email != "" || password != "" {

		key := <-cha
		if key != nil {
			ds := make(chan bool)
			ps := key.Map()["subdata"].(primitive.D).Map()["password"].(string)
			go h.Vcheck(ps, password, ds)
			if <-ds {
				t := time.Now()
				un := key.Map()["username"].(string)
				tag := key.Map()["tag"].(string)
				g, _ := jwt.GenerateToken(c, tag, un, string(t.Format("2006-01-02 15:04:05")), int64(60456))
				strSessionOTP := g
				OTP := GenOTP()
				jwthash := h.Mhash(strSessionOTP)
				Sessionhash := h.Mhash(OTP + "." + string(t.Format("15:04:05")))

				s.Db_FixOneStuck(bson.M{"email": bson.M{"$eq": email}, "username": bson.M{"$eq": un}}, bson.M{"$push": bson.M{"SessionOTP": bson.M{string(t.Format("2006-01-02 15:04:05")): jwthash + " " + Sessionhash}}})
				am.SEndlogin(un, tag, OTP, email)
				c.JSON(200, gin.H{
					"message": "login suss",
					"jwt":     g,
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
