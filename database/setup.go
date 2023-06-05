package database

import (
	"hubla-challenge/database/models"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

func Setup() {
	db, err := gorm.Open("postgres", "host=localhost user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	models.Setup(db)
}
