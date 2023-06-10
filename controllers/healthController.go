package controllers

import (
	"hubla-challenge/interfaces"
	"hubla-challenge/services"

	"github.com/gofiber/fiber/v2"
)

type HealthController struct{}

func (*HealthController) Route(app *fiber.App) {
	app.Get("/", services.HealthCheck)
}

func NewHealthController() interfaces.Router {
	healthController := &HealthController{}

	return healthController
}
