package verifyuser

import (
	db "api/db"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenKEy() string {

	bytes := make([]byte, 19)
	var pool = "234567890abcdefghijxk2345678lmopqrsx2345678yzabxcdxefghxijkl2345678mopqrsyzx???>><<|\\}{!#@$%^&*())__"
	for i := 0; i < 19; i++ {

		bytes[i] = pool[rand.Intn(len(string(pool)))]

	}

	return string(bytes)

}
func SaveDAta(s db.Db_mongo, email, password, user, tag, time, id, OBJ string) {
	var post DATA
	post.Time = time
	post.Email = email
	post.Subdata.Password = password
	post.Username = user
	post.UserId = id

	post.Tag = tag

	s.Db_InsertOneS(post)
	s.Db_FixOneStuck(bson.M{"email": bson.M{"$eq": email}, "tag": bson.M{"$eq": tag}, "username": bson.M{"$eq": user}}, bson.M{"$push": bson.M{"SessionAuthor": OBJ}})

}
