package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/utils"
	"gorm.io/driver/postgres"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup(app *fiber.App) {

	dbuser := utils.GetEnv("DBUSER", "admin")
	dbpass := utils.GetEnv("DBPASS", "admin")
	dbhost := utils.GetEnv("DBHOST", "localhost")
	dbname := utils.GetEnv("DBNAME", "hubla")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		dbhost, dbuser, dbname, dbpass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	models.Setup(db)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

}
