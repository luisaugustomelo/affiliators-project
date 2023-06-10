package database

import (
	"hubla-challenge/database/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func Setup() {
	db, err := gorm.Open("postgres", "host=172.22.0.2 user=admin dbname=hubla password=admin sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	models.Setup(db)
}
