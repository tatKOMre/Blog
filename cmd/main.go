package main

import (
	"log"
	"tatKOM/app"
	"tatKOM/model"
	"tatKOM/pkg/db"
)

func main() {

	// подключение к бд
	appDB, err := db.ConnectToDataBase(
		"db",
		"root",
		"5432",
		"admin",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	// миграции моделек в бд
	model.AutoMigrate(appDB)
	log.Println("migration complete")

	blogApp := app.New(appDB, []byte("niggaballs"), ":8080")
	log.Println("server initialized")

	// запуск сервера
	blogApp.Run()
}
