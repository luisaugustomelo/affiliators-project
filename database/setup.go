package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Setup(app *fiber.App) {

	dbuser := utils.GetEnv("DBUSER", "admin")
	dbpass := utils.GetEnv("DBPASS", "admin")
	dbhost := utils.GetEnv("DBHOST", "localhost")
	dbname := utils.GetEnv("DBNAME", "hubla")

	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
			dbhost, dbuser, dbname, dbpass),
	)

	if err != nil {
		panic("failed to connect database")
	}

	models.Setup(db)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

}
