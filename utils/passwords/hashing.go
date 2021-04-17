package passwords

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSaltPassword(password *string) {
	bytePass := []byte(*password)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.MinCost)
	if err != nil {
		log.Println(err.Error())
	}
	*password = string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	plainBytes := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainBytes)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
