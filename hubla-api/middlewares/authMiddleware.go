package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/services"
)

func RenewJWT(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(services.JwtSecret), nil
	})

	if err != nil {
		// Handle error parsing the token
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims := t.Claims.(jwt.MapClaims)

	// Verify the token hasn't expired
	expirationTimeUnix := claims["exp"].(float64) // JWT lib uses float64 to represent the date
	currentTimeUnix := time.Now().Unix()
	if currentTimeUnix > int64(expirationTimeUnix) {
		// Token has expired, reject the renewal
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token expired",
		})
	}

	email := claims["sub"].(string)
	id := claims["id"].(float64)

	newToken, err := services.GenerateJWT(email, uint(id))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Set the new token in the response header
	c.Set("X-New-Token", newToken)
	c.Locals("user", interfaces.Credentials{
		Email: email,
		Id:    uint(id),
	})

	return c.Next()
}
