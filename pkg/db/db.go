package db

import (
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDataBase(dbName, dbPass, dbPort, dbUser string) (*gorm.DB, error) {
	dsn := []string{
		"dbname=" + dbName,
		"port=" + dbPort,
		"user=" + dbUser,
		"password=" + dbPass,
		"host=PostSQL",
		"sslmode=disable",
	}
	db, err := gorm.Open(postgres.Open(strings.Join(dsn, " ")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	},
	)
	return db, err
}
