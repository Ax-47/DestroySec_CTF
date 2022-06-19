package verify

import (

	//"fmt"
	"math/rand"
)

func GenKEy() string {

	bytes := make([]byte, 19)
	var pool = "234567890abcdefghijxk2345678lmopqrsx2345678yzabxcdxefghxijkl2345678mopqrsyzx???>><<|\\}{!#@$%^&*())__"
	for i := 0; i < 19; i++ {

		bytes[i] = pool[rand.Intn(len(string(pool)))]
	}
	return string(bytes)

}
