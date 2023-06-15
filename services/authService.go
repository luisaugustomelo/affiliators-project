package services

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("secret")

func GenerateJWT(id string) (string, error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		Id:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RenewJWT(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	claims := t.Claims.(jwt.MapClaims)
	idValue := claims["jti"].(string)

	newToken, err := GenerateJWT(idValue)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Set the new token in the response header
	c.Set("X-New-Token", newToken)

	return c.Next()
}
