package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3000"
	}

	app.Listen(PORT)
}
