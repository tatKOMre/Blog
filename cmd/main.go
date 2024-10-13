package main

import (
	"log"
	"tatKOM/app"
	"tatKOM/pkg/db"
)

func main() {

	parseFlags()

	// подключение к бд
	appDB, err := db.Connect(
		"имя db",
		"порт db",
		"логин db",
		"пароль db",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	// миграции моделек в бд
	model.Migrate(appDB)
	log.Println("migration complete")

	blogApp := app.New(appDB, []byte("niggaballs"), ":8080")
	log.Println("server initialized")

	// запуск сервера
	usersApp.Run()
}
