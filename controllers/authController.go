package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"github.com/luisaugustomelo/hubla-challenge/services"
	"gorm.io/gorm"
)

type AuthController struct{}

func Auth(c *fiber.Ctx) error {
	user := new(models.User)
	db := c.Locals("db").(*gorm.DB)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	userExists, err := services.GetUserByEmail(db, user.Email)
	if err != nil {
		return c.JSON(err)
	}

	if userExists.Password != user.Password {
		return c.JSON("user not authorized")
	}

	jwt, err := services.GenerateJWT(user.Email, userExists.ID)

	if err != nil {
		return c.JSON(err)
	}

	err = services.UpdateUser(db, userExists.ID, userExists)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(map[string]interface{}{
		"token": jwt,
		"user":  userExists.ID,
	})
}

func (*AuthController) Route(app *fiber.App) {
	app.Post("/login", Auth)
}

func NewAuthController() interfaces.Router {
	authController := &AuthController{}

	return authController
}
