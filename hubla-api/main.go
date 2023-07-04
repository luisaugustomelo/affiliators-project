package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/luisaugustomelo/hubla-challenge/controllers"
	"github.com/luisaugustomelo/hubla-challenge/database"
	"github.com/luisaugustomelo/hubla-challenge/workers"
)

var app *fiber.App

func init() {
	app = fiber.New()
	app.Use(cors.New())

	db, err := database.Setup(app)
	if err != nil {
		panic(err)
	}
	workers.ConsumerToQueue(db)
}

func main() {
	controllers.SetupRoutes(app)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3030"
	}

	app.Listen(PORT)
}
