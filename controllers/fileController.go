package controllers

import (
	"fmt"

	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/services"

	"github.com/gofiber/fiber/v2"
)

type FileController struct{}

func fiberError(c *fiber.Ctx, status int, message string, err error) error {
	return c.Status(status).SendString(fmt.Sprintf("%s : %s", message, err.Error()))
}

func UploadSingleFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fiberError(c, fiber.StatusBadRequest, "File error", err)
	}

	filename, err := services.ProcessFile(file)
	if err != nil {
		return fiberError(c, fiber.StatusInternalServerError, "Failed to process file", err)
	}

	transactions, err := services.ReadTransactions(filename)
	if err != nil {
		return fiberError(c, fiber.StatusInternalServerError, "Failed to read transactions", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"filepath": "/images/single/" + filename, "transactions": transactions})
}

func (*FileController) Route(app *fiber.App) {
	app.Post("/upload", UploadSingleFile)
}

func NewFileController() interfaces.Router {
	fileController := &FileController{}

	return fileController
}
