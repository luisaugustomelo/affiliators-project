package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/middlewares"
	"github.com/luisaugustomelo/hubla-challenge/services"
	"github.com/luisaugustomelo/hubla-challenge/workers"
	"gorm.io/gorm"
)

type FileController struct{}

func fiberError(c *fiber.Ctx, status int, message string, err error) error {
	return c.Status(status).SendString(fmt.Sprintf("%s : %s", message, err.Error()))
}

func CheckFileStatus(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	value, err := strconv.Atoi(id)

	if err != nil {
		return fiberError(c, fiber.StatusBadRequest, "ID error", fmt.Errorf("id doesn't exist"))
	}

	balances, err := services.CheckFileStatus(db, value)
	if err != nil {
		return fiberError(c, fiber.StatusBadRequest, "ID process error", fmt.Errorf("error to process request"))
	}

	return c.Status(200).JSON(balances)
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
	user := c.Locals("user").(interfaces.Credentials)
	mq, err := workers.PublishToQueue(interfaces.Message{
		UserId: user.Id,
		Email:  user.Email,
		File:   base64.StdEncoding.EncodeToString(bytes),
	}, db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(mq)
}

func (*FileController) Route(app *fiber.App) {
	app.Post("/upload", middlewares.RenewJWT, UploadSingleFile)
	app.Get("/checkStatus/:id", CheckFileStatus)
}

func NewFileController() interfaces.Router {
	fileController := &FileController{}

	return fileController
}
