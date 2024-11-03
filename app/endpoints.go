package app

import (
	"crypto/tls"
	"log"
	"net/http"
	"tatKOM/pkg/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
)

// Метод для привязки функция к адресам на сайте
func (app *App) CreateEndpoints() {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMW)

	// Это для подгрузки css/js файлов
	r.PathPrefix("/web/").Handler(
		http.StripPrefix(
			"/web",
			http.FileServer(http.Dir("./web/")),
		),
	)

	//Это чтобы js не выебывался, просто не трогай, сам не ебу че там
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}
	server := &http.Server{
		Addr:    ":443",
		Handler: r,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	server.ListenAndServeTLS("", "")
	log.Println("certificates  are ready")

	handler := c.Handler(r)
	r.HandleFunc("/", app.Handler.CheckUserPermission)
	r.HandleFunc("/publications/", app.Handler.GetAllPublications)
	r.HandleFunc("/login/", app.Handler.Signin)
	r.HandleFunc("/registration/", app.Handler.SignUp)
	r.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/html/admin.html")
	})
	r.HandleFunc("/profile/", app.Handler.CheckUserPermissionProfile)
	r.HandleFunc("/admin/crpub", app.MW.Auth(app.Handler.CreatePublication))
	r.HandleFunc("/admin/delpub", app.MW.Auth(app.Handler.DeletePublication))

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

		Если только для залогиненных, то:
		r.HandleFunc("/nigga/", app.MW.Auth(app.Handler.Nigga))
	*/
	(*app).Server.Handler = handler
}
