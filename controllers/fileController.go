package controllers

import (
	"hubla-challenge/services"

	"github.com/gofiber/fiber/v2"
)

type FileController struct{}

func (*FileController) Route(app *fiber.App) {
	app.Post("/upload", services.UploadSingleFile)
}
