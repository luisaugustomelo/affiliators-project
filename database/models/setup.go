package models

import (
	"gorm.io/gorm"
)

var models = []interface{}{
	&User{},
	&Product{},
	&Affiliate{},
	&Creator{},
	// Adicione mais modelos aqui conforme necessário
}

func Setup(db *gorm.DB) {
	for _, model := range models {
		db.AutoMigrate(model)
	}
}
