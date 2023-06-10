package main

import (
	"os"

	"github.com/luisaugustomelo/hubla-challenge/controllers"
	"github.com/luisaugustomelo/hubla-challenge/database"

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
