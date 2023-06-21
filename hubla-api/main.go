package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/luisaugustomelo/hubla-challenge/controllers"
	"github.com/luisaugustomelo/hubla-challenge/database"
	"github.com/luisaugustomelo/hubla-challenge/workers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	db, err := database.Setup(app)
	if err != nil {
		panic(err)
	}
	time.Sleep(500)
	workers.ConsumerToQueue(db)

	controllers.SetupRoutes(app)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3030"
	}

	app.Listen(PORT)
}
