package controllers

import (
	"hubla-challenge/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")

	api.Get("/", services.HelloWorld)

	api.Post("/upload", services.UploadSingleFile)
}
