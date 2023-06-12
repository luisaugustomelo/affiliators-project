package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/services"

	"github.com/luisaugustomelo/hubla-challenge/interfaces"
)

type UserController struct{}

// CreateUser is controller to create a new user.
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

// GetUser is controller to get an user.
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

// UpdateUser is controller to update an user.
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

// DeleteUser is controller to delete an user.
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

func (u *UserController) Route(app *fiber.App) {
	app.Get("/users/:id", GetUser)
	app.Post("/create", CreateUser)
	app.Put("/update", UpdateUser)
	app.Delete("/users/:id", DeleteUser)
}

func NewUserController() interfaces.Router {
	userController := &UserController{}

	return userController
}
