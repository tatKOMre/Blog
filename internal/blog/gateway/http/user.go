package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"tatKOM/model"
	"tatKOM/pkg/token"

	"github.com/gorilla/mux"
)

/*
Тут функции логина/регистрации надо ебануть, это тебе известно,
При логине клади пользователю в куки файлы токен,
потом через middleware для авторизации будешь его получать и проверять
В pkg/cookie функции для этого есть
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
