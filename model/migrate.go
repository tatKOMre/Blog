package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Publication{})
	db.AutoMigrate(&User{})
}
