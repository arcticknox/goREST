package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(plaintext []byte) (string, error) {
	password, error := bcrypt.GenerateFromPassword(plaintext, bcrypt.DefaultCost)
	if error != nil {
		return "", error
	}
	return string(password), nil
}

func ComparePasswordHash(hashedPassword string, password string) error {
	checkPass := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return checkPass
}
