package models

import (
	"gorm.io/gorm"
)

type QueueProcessing struct {
	gorm.Model
	UserId  uint `gorm:"foreignKey:user_id"`
	Status  string
	Message string
	Hash    string
}
