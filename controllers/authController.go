package controllers

import (
	"github.com/luisaugustomelo/hubla-challenge/interfaces"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct{}

func (*AuthController) Route(app *fiber.App) {
	//app.Get("/users/:id", nil)
	//app.Post("/login", nil)
	//app.Post("/logout", nil)
	//app.Delete("/users/:id", nil)
}

func NewAuthController() interfaces.Router {
	authController := &AuthController{}

	return authController
}
