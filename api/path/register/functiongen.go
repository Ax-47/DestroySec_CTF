package register

import (
	db "api/db"

	"math/rand"
)

func GanuserTag(s db.Db_mongo) string {
	for {
		bytes := make([]byte, 4)
		var pool = "1234567890"
		for i := 0; i < 4; i++ {
			bytes[i] = pool[rand.Intn(len(pool))]
		}
		some, _ := s.Db_FindALL("userid", string(bytes))
		if some == nil {
			return string(bytes)
		}
	}

}
func Ganuserid(s db.Db_mongo) string {
	for {
		var pool = "1234567890"
		dd := make([]byte, 13)
		for i := 0; i < 13; i++ {
			dd[i] = pool[rand.Intn(len(pool))]
		}
		some, _ := s.Db_FindALL("userid", string(dd))
		if some == nil {
			return string(dd)
		}
	}

}
