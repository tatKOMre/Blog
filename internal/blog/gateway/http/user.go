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
 Продумай какие нас сайте будут страницы, продуймай какие для каждой из них нужны данные,
 Например, для страницы с постом понадобится запись поста из бд и все комментарии к нему.
 Тоесть, в gateway функции этой страницы понадобится вызывать service.GetPost, service.GetCommentsFor
 И так со всеми страницами на сайте, пока что можешь таким образом переделать, потом с папкой app/ разберемся
*/

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "request", r)
	if r.Method == http.MethodPost {
		type loginReq struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}

		var req loginReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user := model.User{
			Login:      req.Login,
			Password:   req.Password,
			Permission: true,
		}

		err := h.Service.SignUp(ctx, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	} else if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/html/register.html")
	}
}

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "request", r)
	if r.Method == http.MethodPost {
		type log struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		var l log
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tkn, err := h.Service.Signin(ctx, l.Login, l.Password)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		print(tkn)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": tkn})
	} else if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/html/login.html")
	}
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
