package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte("secret")

func GenerateJWT(id string) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 60).Unix(),
		"sub": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
