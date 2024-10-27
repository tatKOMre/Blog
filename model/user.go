package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	Login      string `gorm:"not null"`
	Password   string `gorm:"not null"`
	Permission bool
}
