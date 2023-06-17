package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Email           string `gorm:"type:varchar(100);unique;not null"`
	Password        string `json:"-"`
	Role            uint   `gorm:"foreignKey:role"`
	QueueProcessing QueueProcessing
	SalesFile       SalesFile `gorm:"foreignKey:UserID"`
}
