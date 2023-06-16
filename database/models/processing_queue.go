package models

import (
	"gorm.io/gorm"
)

type ProcessingQueue struct {
	gorm.Model
	Status string
	UserId uint
	Hash   string
}
