package login

import (
	db "api/db"
	h "api/hashpaww"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ln struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context, s db.Db_mongo) {

	var fromreg ln
	if err := c.BindJSON(&fromreg); err != nil {
		return
	}
	email := fromreg.Email
	password := fromreg.Password
	if email != "" || password != "" {
		ch := make(chan primitive.D)
		go s.Db_FindtOne("email", email, ch)
		key := <-ch
		if key != nil {

			if h.Vcheck(key.Map()["subdata"].(primitive.D).Map()["password"].(string), password) {

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
