package crypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Encrypter interface {
	EncryptPassword(p string) string
}

type BCrypt struct {
}

// Returns password encrypted to insert in the database
func (b *BCrypt) EncryptPassword(p string) string {
	p_bytes := []byte(p)
	password_bytes, err := bcrypt.GenerateFromPassword(p_bytes, bcrypt.MinCost)
	if err != nil {
		log.Fatalln("Failed to generate password hash")
	}
	new_password := string(password_bytes)
	return new_password
}
