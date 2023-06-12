package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/services"
)

type UserController struct{}

// CreateUser é o controlador para criar um novo usuário.
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	db := c.Locals("db").(*gorm.DB)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := services.CreateUser(db, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser é o controlador para obter um usuário.
func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	db := c.Locals("db").(*gorm.DB)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}
	user, err := services.GetUser(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}
	return c.JSON(user)
}

// UpdateUser é o controlador para atualizar um usuário.
func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)
	db := c.Locals("db").(*gorm.DB)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := services.UpdateUser(db, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}
	return c.JSON(user)
}

// DeleteUser é o controlador para deletar um usuário.
func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	db := c.Locals("db").(*gorm.DB)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}
	if err := services.DeleteUser(db, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "User deleted",
	})
}

func (*UserController) Route(app *fiber.App) {
	app.Get("/users/:id", GetUser)
	app.Post("/create", CreateUser)
	app.Put("/update", UpdateUser)
	app.Delete("/users/:id", DeleteUser)
}

func NewUserController() interfaces.Router {
	userController := &UserController{}

	return userController
}
