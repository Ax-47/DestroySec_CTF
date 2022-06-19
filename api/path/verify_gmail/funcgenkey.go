package verify

import (
	db "api/db"
	//"fmt"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenKEy(s db.Db_mongo, datauser primitive.D) string {

	bytes := make([]byte, 19)
	var pool = "234567890abcdefghijxk2345678lmopqrsx2345678yzabxcdxefghxijkl2345678mopqrsyzx???>><<|\\}{!#@$%^&*())__"
	for i := 0; i < 19; i++ {

		bytes[i] = pool[rand.Intn(len(string(pool)))]
	}
	return string(bytes)

}
