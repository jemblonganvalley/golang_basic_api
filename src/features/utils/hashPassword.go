package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string)(string, error){
	hashResult,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if(err != nil){
		log.Println(err)
		return "", err
	}
	return string(hashResult), nil 
}