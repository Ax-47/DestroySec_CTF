package hashpaww

import (
	"golang.org/x/crypto/bcrypt"
)

func Mhash(passworsd string) string {
	password := []byte(passworsd)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func Vcheck(word, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(word), []byte(password))
	//fmt.Println(err)
	return err == nil
}
