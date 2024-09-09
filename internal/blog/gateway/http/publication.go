package http

import (
	"context"
	"net/http"
	"strconv"
	"tatKOM/model"
	"tatKOM/pkg/token"

	"github.com/gorilla/mux"
)

func (h *Handler) CreatePublication(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	ctx := context.WithValue(context.Background(), "request", r)
	var publication model.Publication
	err := h.Service.CreatePublication(ctx, publication, act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) DeletePublication(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctx := context.WithValue(context.Background(), "request", r)
	err := h.Service.DeletePublication(ctx, uint(id), act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
