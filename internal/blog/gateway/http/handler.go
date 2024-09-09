package http

import (
	"tatKOM/internal/blog"
)

type Handler struct {
	Service blog.Service
	SignKey []byte
}

func New(service blog.Service, key []byte) *Handler {
	return &Handler{
		Service: service,
		SignKey: key,
	}
}
