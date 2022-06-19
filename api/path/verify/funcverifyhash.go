package Check

import (
	h "api/hash_class"
)

func Asyncverify(a, b string, stat chan Checkverify) {
	C := make(chan bool)
	go h.Vcheck(a, b, C)
	if <-C == true {
		stat <- Checkverify{true}
	} else {
		stat <- Checkverify{false}
	}
}
