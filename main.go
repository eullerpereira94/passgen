package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "senha"

	hash, _ := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	fmt.Println(string(hash))
}
