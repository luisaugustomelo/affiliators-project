package models

import (
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	Type    string
	UserId  uint `gorm:"type:integer;not null"`
	Date    string
	Product string
	Value   string
	Seller  string
	Hash    string
}
