package http

import (
	"context"
	"net/http"
	"tatKOM/model"
	"tatKOM/pkg/token"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	ctx := context.WithValue(context.Background(), "request", r)
	var comment model.Comment
	err := h.Service.CreateComment(ctx, comment, act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	ctx := context.WithValue(context.Background(), "request", r)
	err := h.Service.DeleteComment(ctx, act.ID, act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request, act *token.Claims) {
	ctx := context.WithValue(context.Background(), "request", r)
	comment, err := h.Service.GetComment(ctx, act.ID, act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	err = h.Service.UpdateComment(ctx, comment, act)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

if r.Method == http.MethodGet {

}
if r.Method == http.MethodPost {
	r.ParseForm()

	title := r.FormValue("title")
}