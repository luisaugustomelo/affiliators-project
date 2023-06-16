package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/luisaugustomelo/hubla-challenge/controllers"
	"github.com/luisaugustomelo/hubla-challenge/database"
	"github.com/luisaugustomelo/hubla-challenge/workers"
)

func main() {
	app := fiber.New()
	database.Setup(app)

	workers.ConsumerToQueue()

	controllers.SetupRoutes(app)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3030"
	}

	app.Listen(PORT)
}
