package model

import (
	"gorm.io/gorm"
)

type Publication struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null" json:"title"`
	Text  string `gorm:"not null" json: "text"`
	Views uint   `gorm:"default:0"`
}
