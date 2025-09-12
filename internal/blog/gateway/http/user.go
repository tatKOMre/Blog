package http

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"tatKOM/model"
	"tatKOM/pkg/cookie"
	"tatKOM/pkg/token"

	"github.com/gorilla/mux"
)


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
			Permission: false,
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
func (h *Handler) CheckUserPermission(w http.ResponseWriter, r *http.Request) {
	usr := &token.Claims{}
	cookie, err := cookie.GetCookie(r, "token")
	if err == nil {
		usr, _ = token.ParseJWT(cookie, h.SignKey)
	}

	tmpl := template.Must(template.ParseFiles("./web/html/index.html"))

	tmpl.Execute(w, usr)
}
func (h *Handler) CheckUserPermissionProfile(w http.ResponseWriter, r *http.Request) {
	usr := &token.Claims{}
	cookie, err := cookie.GetCookie(r, "token")
	if err == nil {
		usr, _ = token.ParseJWT(cookie, h.SignKey)
	}

	tmpl := template.Must(template.ParseFiles("./web/html/profile.html"))

	tmpl.Execute(w, usr)
}
