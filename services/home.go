package services

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
