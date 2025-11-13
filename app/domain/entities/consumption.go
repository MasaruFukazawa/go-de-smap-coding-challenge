package entities

import (
	"gorm.io/gorm"
	"time"
)

type Consumption struct {
	gorm.Model
	User     User `gorm:"foreignKey:UserID"`
	Datetime time.Time
	Value    float64 `gorm:"type:decimal(10,2)"`
}
