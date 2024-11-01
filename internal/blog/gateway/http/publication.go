package http

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"tatKOM/model"
	"tatKOM/pkg/token"

	"github.com/gorilla/mux"
)

func (h *Handler) CreatePublication(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	if r.Method == http.MethodPost {
		var pub model.Publication
		if err := json.NewDecoder(r.Body).Decode(&pub); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(context.Background(), "request", r)
		err := h.Service.CreatePublication(ctx, pub, act)

		if err != nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) GetAllPublications(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "request", r)

	publications, err := h.Service.GetAllPublications(ctx)

	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/html/publication.html"))
	tmpl.Execute(w, publications)
}

func (h *Handler) DeletePublication(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	if r.Method == http.MethodPost {
		type reqData struct {
			ID uint `json:"id"`
		}
		var id reqData
		print("312123123123123123123123123")
		err := json.NewDecoder(r.Body).Decode(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		print("sdaasdasdsdaasd")
		ctx := context.WithValue(context.Background(), "request", r)
		err = h.Service.DeletePublication(ctx, id.ID, act)

		if err != nil {
			http.Error(w, "server error", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) UpdatePublication(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctx := context.WithValue(context.Background(), "request", r)
	publication, err := h.Service.GetPublication(ctx, uint(id), act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	er := h.Service.UpdatePublication(ctx, publication, act)
	if er != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
