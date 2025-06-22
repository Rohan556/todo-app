package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	const cost = 10

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(encryptedPassword), err
}
