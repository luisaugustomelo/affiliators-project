package models

import (
	"gorm.io/gorm"
)

type SalesFile struct {
	gorm.Model
	Type      string
	UserID    uint `gorm:"foreignKey:user_id"`
	SalesType uint `gorm:"foreignKey:sales_type"`
	Date      string
	Product   string
	Value     string
	Seller    string
	Hash      string
}
