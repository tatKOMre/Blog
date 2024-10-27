package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	Login      string `gorm:"not null" json:"login"`
	Password   string `gorm:"not null" json:"password"`
	Permission bool
}
