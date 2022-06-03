package login

import (
	db "api/db"
	h "api/hash_class"
	jwt "api/jwt/service"
	"fmt"

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
				g, _ := jwt.GenerateToken(c, un, "1", int64(60456))
				fmt.Println(g)
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
