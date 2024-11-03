package app

import (
	"log"
	"net/http"
	handler "tatKOM/internal/blog/gateway/http"
	"tatKOM/internal/blog/repository"
	"tatKOM/internal/blog/service"
	"tatKOM/pkg/middleware"

	"github.com/gin-gonic/autotls"
	"golang.org/x/crypto/acme/autocert"
	"gorm.io/gorm"
)

// App - структура
type App struct {
	Handler *handler.Handler
	Router  http.Handler
	MW      *middleware.Middleware
	SignKey []byte
	Manager autocert.Manager
	DB      *gorm.DB
}

// Run - Эта функция запускает сервер
func (app *App) Run() {
	log.Println("certificates  are ready")

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("protatkom.ru"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Println("server running")
	log.Fatal(autotls.RunWithManager(app.Router, &m))

}

// New - создание структуры App
func New(db *gorm.DB, key []byte) *App {
	app := &App{
		DB:      db,
		SignKey: key,
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
