package main

import (
	"os"

	"hubla-challenge/controllers"
	"hubla-challenge/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	controllers.SetupRoutes(app)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3030"
	}

	database.Setup()

	app.Listen(PORT)
}
