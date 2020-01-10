package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password :="fangkaiming"
	hashpass :="$2a$10$3Hkx7h/Q54GwiVMffWZV1.ah6EIQQ3MJZcK1m0tmGb6Ahrz8liZdq"


	//hashPass,_ :=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	//fmt.Println(string(hashPass))

	err :=bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password))
	if err !=nil{
		fmt.Println(err)
	}
}
