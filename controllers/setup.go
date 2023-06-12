package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//api := app.Group("/")

	authController := NewAuthController()
	authController.Route(app)

	healthController := NewHealthController()
	healthController.Route(app)

	fileController := NewFileController()
	fileController.Route(app)

	userController := NewUserController()
	userController.Route(app)
}
