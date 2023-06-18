package models

import (
	"gorm.io/gorm"
)

type SalesFile struct {
	gorm.Model
	UserID    uint `gorm:"foreignKey:user_id"`
	SalesType uint `gorm:"foreignKey:sales_type"`
	Date      string
	Product   string
	Value     float64
	Seller    string
	Hash      string
}
