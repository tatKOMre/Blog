package app

import (
	"tatKOM/pkg/middleware"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

// Метод для привязки функция к адресам на сайте
func (app *App) Route() {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMW)

	// Это для подгрузки css/js файлов
	r.PathPrefix("/web/").Handler(
		http.StripPrefix(
			"/web",
			http.FileServer(http.Dir("./web/")),
		),
	)

	// Это чтобы js не выебывался, просто не трогай, сам не ебу че там
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	// Вот тут Привязка
	/*
	r.HandleFunc("/login/", app.Handler.Login).
		Methods(http.MethodPost)
	Такая хуйня означает, 
	что функция логин вызовется только при POST запросе по адресу <>.ru/login/

	r.HandleFunc("/fuck/", app.Handler.Nigger)
	Эта хуйня будет вызывать функцию при запросе по адресу <>.ru/fuck/ с любым методом

	И таким образом ты подвязываешь каждую функцию из gateway к адресу на сайте,
	к одному адресу модно подвязать несколько функций, если каждая из них обрабатывает свои методы запроса
	*/
	app.Server.Handler = handler
}
