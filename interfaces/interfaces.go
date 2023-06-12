package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Router interface {
	Route(app *fiber.App)
}

type Datastore interface {
	Create(value interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
}
