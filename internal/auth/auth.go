package auth

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func EncryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(encryptedPassword), err
}

func IsPasswordValid(hashedPassword, passwordInput string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordInput))

	if err != nil {
		return false
	}

	return true
}
