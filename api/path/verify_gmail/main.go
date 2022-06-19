package verify

import (
	db "api/db"
	jw "api/jwt/service"

	h "api/hash_class"
	tl "api/timelogic"
	wa "api/when_after_login"
	"fmt"
	_ "fmt"
	"strings"
	"time"

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
	timestap := a.Claims.(jwt.MapClaims)["sub"].(string)
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
	some, _ := s.Db_FindALLD("username", "tag", tag, un)
	ax := some[0].Map()["SessionOTP"].(primitive.A)
	if tl.Logic_time_lteq(timestap, S) {
		for num, volume := range ax {
			if volume != nil {
				ps := volume.(primitive.D).Map()
				hashJwt := ps[timestap]
				if hashJwt != nil {

					ds := make(chan bool)
					//[timestap].(string)
					sd := make(chan bool)

					ax := strings.Split(hashJwt.(string), " ")
					timestapspilt := strings.Split(timestap, " ")
					go h.Vcheck(ax[0], Ax.Jwt, ds)
					go h.Vcheck(ax[1], Ax.OTP+"."+timestapspilt[1], sd)
					if <-ds && <-sd {

						//fmt.Print(ps)

						//s.Db_FixOneStuck(bson.M{"tag": bson.M{"$eq": tag}, "username": bson.M{"$eq": un}}, bson.M{"$unset": bson.M{"SessionOTP": bson.A{bson.M{timestap: Ax.Jwt}}}})

						OBJ := GenKEy()
						hash, _ := jw.GenerateTokenNEW(c, tag, un, OBJ)
						hashJWt := h.Mhash(hash)
						//SaveDAta(s, email, password, un, tag, ax["time"].(string), ax["userid"].(string), hashJWt)
						s.Db_FixOneStuck(bson.M{"tag": bson.M{"$eq": tag}, "username": bson.M{"$eq": un}}, bson.M{"$push": bson.M{"SessionAuthor": hashJWt}})
						s.Db_FixOneStuck(bson.M{"tag": bson.M{"$eq": tag}, "username": bson.M{"$eq": un}}, bson.M{"$unset": bson.M{"SessionOTP.$[]": num}})

						c.JSON(200, gin.H{
							"message":        "login suss",
							"djkfhjdhgfjdfd": wa.Index(c, tag, un, OBJ),
						})
					} else {
						c.JSON(405, gin.H{
							"message": "login suss",
						})
					}

					break
				}
			}

		}
	} else {
		c.JSON(403, gin.H{
			"message": "login suss",
		})
	}

}
