package app

import (
	"crypto/tls"
	"log"
	"net/http"
	handler "tatKOM/internal/blog/gateway/http"
	"tatKOM/internal/blog/repository"
	"tatKOM/internal/blog/service"
	"tatKOM/pkg/middleware"

	"golang.org/x/crypto/acme/autocert"
	"gorm.io/gorm"
)

// App - структура
type App struct {
	Handler *handler.Handler
	Server  *http.Server
	MW      *middleware.Middleware
	SignKey []byte
	Manager autocert.Manager
	DB      *gorm.DB
}

// Run - Эта функция запускает сервер
func (app *App) Run() {
	log.Println("server running")

	if err := app.Server.ListenAndServe(); err != nil {
		log.Println(err)
	}

	go http.ListenAndServe(":80", app.Manager.HTTPHandler(nil))
	app.Server.ListenAndServeTLS("", "")
	log.Println("certificates  are ready")
}

// New - создание структуры App
func New(db *gorm.DB, key []byte) *App {
	app := &App{
		DB:      db,
		SignKey: key,
	}

	app.Manager = autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}

	app.Server = &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			GetCertificate: app.Manager.GetCertificate,
		},
	}

	app.MW = middleware.New(app.SignKey)

	// Все слои создаем
	r := repository.New(db)
	s := service.New(r, key)
	h := handler.New(s, key)

	app.Handler = h

	app.CreateEndpoints()

	return app
}
