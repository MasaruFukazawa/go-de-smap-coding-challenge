package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID uint
	Area   uint
	Tariff uint
}
