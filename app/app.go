package app

import (
	"log"
	"net/http"
	"os"
	handler "tatKOM/internal/blog/gateway/http"
	"tatKOM/internal/blog/repository"
	"tatKOM/internal/blog/service"
	"tatKOM/pkg/middleware"
	"time"

	"gorm.io/gorm"
)

// App - структура
type App struct {
	Handler *handler.Handler
	Server  *http.Server
	MW      *middleware.Middleware
	SignKey []byte
	DB      *gorm.DB
}

// Run - Эта функция запускает сервер
func (app *App) Run() {
	log.Println("server running")

	if err := app.Server.ListenAndServe(); err != nil {
		log.Println(err)
	}

	log.Println("shutting down")
	os.Exit(0)
}

// New - создание структуры App
func New(db *gorm.DB, key []byte, addr string) *App {
	app := &App{
		DB:      db,
		SignKey: key,
	}

	app.Server = &http.Server{
		Addr:         addr,             // порт
		WriteTimeout: 15 * time.Second, // таймауты
		ReadTimeout:  15 * time.Second,
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
