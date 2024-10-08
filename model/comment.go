package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Text          string `gorm:"not null"`
	UserID        uint `gorm:"not null"` 
	PublicationID uint `gorm:"not null"`
}
