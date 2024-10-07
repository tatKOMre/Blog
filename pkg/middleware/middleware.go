package middleware

import (
	"net/http"
	"tatKOM/pkg/cookie"
	"tatKOM/pkg/token"
)

type Middleware struct {
	Signkey []byte
}

func New(key []byte) *Middleware {
	return &Middleware{
		Signkey: key,
	}
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, act *token.Claims)

func (m *Middleware) Auth(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := cookie.GetCookie(r, "login")
		tkn, err := token.ParseJWT(cookie, m.Signkey)
		if err != nil {
			http.Error(w, "auth error", http.StatusForbidden)
			return
		}
		f(w, r, tkn)
	}
}

func (m *Middleware) NotAuth(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := cookie.GetCookie(r, "login")
		tkn, err := token.ParseJWT(cookie, m.Signkey)
		if err != nil {
			f(w, r, tkn)
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

/*
Добавь middleware для авторизации

type HandlerFunc func(w http.ResponseWriter, r *http.Request, act *token.Claims)

func (m *Middleware) Auth(f HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Тут получение токена из кукисов
		// и его проверка
		f(w, r, act)
	}
}
*/
