package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte("secret")

func GenerateJWT(email string, id uint) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 60).Unix(),
		"sub": email,
		"id":  id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
