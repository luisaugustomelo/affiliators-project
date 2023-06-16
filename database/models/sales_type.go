package models

import (
	"gorm.io/gorm"
)

type Sale_Type struct {
	gorm.Model
	Description string `gorm:"type:varchar(30);not null"`
	Kind        string `gorm:"type:varchar(30);not null"`
	Signal      string `gorm:"type:varchar(1);not null"`
}
