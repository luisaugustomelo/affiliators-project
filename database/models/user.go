package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID       uint64
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}
