package verifyuser

import (
	db "api/db"
	jw "api/jwt/service"

	tl "api/timelogic"
	"fmt"
	_ "fmt"
	"strings"
	"time"

	h "api/hash_class"
	wa "api/when_after_login"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Verifyotp(c *gin.Context, s db.Db_mongo) {
	t := time.Now()
	var Ax GEtheader
	if err := c.ShouldBindHeader(&Ax); err != nil {
		fmt.Println(err)
	}
	a, err := jw.DecodeToken(Ax.Jwt)
	if err != nil {
		fmt.Println(err)

	}
	//timestap := a.Claims.(jwt.MapClaims)["sub"].(string)
	/*
	   type StandardClaims struct {
	       Audience  string `json:"aud,omitempty"`
	       ExpiresAt int64  `json:"exp,omitempty"`
	       Id        string `json:"jti,omitempty"`
	       IssuedAt  int64  `json:"iat,omitempty"`
	       Issuer    string `json:"iss,omitempty"`
	       NotBefore int64  `json:"nbf,omitempty"`
	       Subject   string `json:"sub,omitempty"`
	   }
	*/

	un := a.Claims.(jwt.MapClaims)["aud"].(string)
	tag := a.Claims.(jwt.MapClaims)["jti"].(string)
	S := string(t.Format("2006-01-02 15:04:05"))
	email := a.Claims.(jwt.MapClaims)["iss"].(string)
	some, _ := s.Db_FindALLD_UniDentify(bson.D{{"email", email}, {"username", un}, {"tag", tag}})

	if some != nil {
		ax := some[0].Map()

		if tl.Logic_time_lteq(ax["time"].(string), S) {

			session := ax["sessionreg"]
			split := strings.Split(session.(string), " ")
			ds := make(chan bool)
			Aha := make(chan bool)

			go h.Vcheck(split[1], Ax.OTP, Aha)
			go h.Vcheck(split[0], Ax.Jwt, ds)
			if <-ds && <-Aha {
				password := ax["subdata"].(primitive.D).Map()["password"].(string)
				OBJ := GenKEy()
				hash, _ := jw.GenerateTokenNEW(c, tag, un, OBJ)
				hashJWt := h.Mhash(hash)
				SaveDAta(s, email, password, un, tag, ax["time"].(string), ax["userid"].(string), hashJWt)
				s.Db_Delete_UniDentify(bson.M{"email": bson.M{"$eq": email}})

				c.JSON(200, gin.H{
					"message":        "login suss",
					"djkfhjdhgfjdfd": wa.Index(c, tag, un, OBJ),
				})
			} else {
				s.Db_Delete_UniDentify(bson.M{"email": bson.M{"$eq": email}})
				c.JSON(403, gin.H{
					"message": "login suss",
				})
			}

		} else {
			s.Db_Delete_UniDentify(bson.M{"email": bson.M{"$eq": email}})
			c.JSON(402, gin.H{
				"message": "time out",
			})
		}

	} else {
		c.JSON(403, gin.H{
			"message": "login suss",
		})
	}
}
