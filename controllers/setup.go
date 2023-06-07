package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//api := app.Group("/")

	authController := AuthController{}
	authController.Route(app)

	healthController := HealthController{}
	healthController.Route(app)

	fileController := FileController{}
	fileController.Route(app)
}
