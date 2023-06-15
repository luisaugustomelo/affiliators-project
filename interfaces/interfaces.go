package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Router interface {
	Route(app *fiber.App)
}

type Datastore interface {
	Create(value interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}
