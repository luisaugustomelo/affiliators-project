package controllers

import (
	"hubla-challenge/services"

	"github.com/gofiber/fiber/v2"
)

type HealthController struct{}

func (*HealthController) Route(app *fiber.App) {
	app.Get("/", services.HealthCheck)
}
