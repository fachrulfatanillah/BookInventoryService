package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my-super-secret-key-1234567890")

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}