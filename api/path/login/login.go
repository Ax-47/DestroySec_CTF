package login

import (
	"api/cookie"
	db "api/db"
	h "api/hashpaww"
	jwt "api/jwt/service"

	//"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *gin.Context, s db.Db_mongo) {

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
				un := key.Map()["username"].(string)
				g, _ := jwt.GenerateToken(c, un, int64(60456))
				Xx := Datacookie{
					user:   un,
					passed: ps,
				}

				cookie.Cookieee_set(c, Xx)
				//cookie.Cookieee_Get(c, []string{"user", "passed"}...)
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
