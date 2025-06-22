package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenInput struct {
	Email string
}

func GenerateJWTToken(email string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")

	fmt.Println(secretKey, "secret")

	claims := jwt.MapClaims{
		"sub": email,
		"iss": "todo-app",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
