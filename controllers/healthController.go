package controllers

import (
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/services"

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
