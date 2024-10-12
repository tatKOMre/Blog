package model

import (
	"gorm.io/gorm"
)

type Publication struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Text  string `gorm:"not null"`
	Views uint   `gorm:"default:0"`
}
