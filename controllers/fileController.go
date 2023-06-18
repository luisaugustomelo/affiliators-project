package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/workers"
	"gorm.io/gorm"
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

	// Check if the file extension is .txt
	if strings.ToLower(filepath.Ext(file.Filename)) != ".txt" {
		return fiberError(c, fiber.StatusBadRequest, "File error", fmt.Errorf("only .txt files are allowed"))
	}

	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	bytes, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return err
	}

	// email will be decrypted based jwt
	db := c.Locals("db").(*gorm.DB)
	workers.PublishToQueue(interfaces.Message{
		UserId: 1,
		Email:  "luis@hubla.com",
		File:   base64.StdEncoding.EncodeToString(bytes),
	}, db)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"filepath": "/images/single/ submited as success!"})
}

func (*FileController) Route(app *fiber.App) {
	app.Post("/upload", UploadSingleFile)
}

func NewFileController() interfaces.Router {
	fileController := &FileController{}

	return fileController
}
