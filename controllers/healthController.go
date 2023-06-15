package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
)

type HealthController struct{}

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Status " + c.Status(200).String())
}

func (*HealthController) Route(app *fiber.App) {
	app.Get("/", HealthCheck)
}

func NewHealthController() interfaces.Router {
	healthController := &HealthController{}

	return healthController
}
