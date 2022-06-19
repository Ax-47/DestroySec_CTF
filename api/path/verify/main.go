package Check

import (
	db "api/db"

	jw "api/jwt/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func subverify(c *gin.Context, s db.Db_mongo) {
	c.JSON(200, gin.H{
		"message": "req fail",
	})
}
func Check(c *gin.Context, s db.Db_mongo) {
	var Ax GETDATA
	if err := c.ShouldBind(&Ax); err != nil {
		fmt.Println(err)
	}
	a, err := jw.DecodeToken(Ax.JWt)
	if err != nil {
		fmt.Println(err)

	}
	tag := a.Claims.(jwt.MapClaims)["jti"].(string)
	name := a.Claims.(jwt.MapClaims)["aud"].(string)
	fmt.Println(a.Claims.(jwt.MapClaims)["aud"].(string))
	fata, _ := s.Db_FindALLD("username", "tag", tag, name)
	ls := fata[0].Map()["SessionAuthor"]
	cha := make(chan Checkverify)
	if fata != nil {
		for _, v := range ls.(primitive.A) {
			go Asyncverify(v.(string), Ax.JWt, cha)
		}
		result := make([]Checkverify, len(ls.(primitive.A)))
		for i, _ := range result {
			result[i] = <-cha
			if result[i].Stats {
				subverify(c, s)
			} else if i == len(ls.(primitive.A)) && !result[i].Stats {
				c.JSON(403, gin.H{
					"message": "req fail",
				})
			}
		}
	} else {
		c.JSON(404, gin.H{
			"message": "req fail",
		})
	}

}
