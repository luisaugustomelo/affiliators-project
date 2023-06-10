package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Route(app *fiber.App)
}
