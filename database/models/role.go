package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Description string
	User        User `gorm:"foreignKey:Role"`
}
