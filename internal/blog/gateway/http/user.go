package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"tatKOM/model"
	"tatKOM/pkg/cookie"
	"tatKOM/pkg/token"

	"github.com/gorilla/mux"
)

/*
 Продумай какие нас сайте будут страницы, продуймай какие для каждой из них нужны данные,
 Например, для страницы с постом понадобится запись поста из бд и все комментарии к нему.
 Тоесть, в gateway функции этой страницы понадобится вызывать service.GetPost, service.GetCommentsFor
 И так со всеми страницами на сайте, пока что можешь таким образом переделать, потом с папкой app/ разберемся
*/

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	ctx := context.WithValue(context.Background(), "request", r)
	var user model.User
	err := h.Service.CreateUser(ctx, user, act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctx := context.WithValue(context.Background(), "request", r)
	user, err := h.Service.GetUser(ctx, uint(id), act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	er := h.Service.UpdateUser(ctx, user, act)
	if er != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctx := context.WithValue(context.Background(), "request", r)
	err := h.Service.DeleteUser(ctx, uint(id), act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctx := context.WithValue(context.Background(), "request", r)
	user, err := h.Service.GetUser(ctx, uint(id), act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// Годно
func (h *Handler) Singin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		login := r.FormValue("login")
		password := r.FormValue("password")
		ctx := context.WithValue(context.Background(), "request", r)
		tkn, err := h.Service.Singin(ctx, login, password)

		if err != nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		cookie.SetCookie(w, "token", tkn)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/html/login.html")
	}
}
