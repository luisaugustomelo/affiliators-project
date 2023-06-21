package models

import (
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	Balance float64
	Role    uint `gorm:"foreignKey:role"`
	Name    string
	Hash    string
}
