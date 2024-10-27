package blog

import (
	"context"
	"tatKOM/model"
	"tatKOM/pkg/token"
)

type Service interface {
	// Взаимодействие с юзером
	GetUser(context.Context, uint, *token.Claims) (model.User, error)
	CreateUser(context.Context, model.User, *token.Claims) error
	UpdateUser(context.Context, model.User, *token.Claims) error
	DeleteUser(context.Context, uint, *token.Claims) error
	Signin(context.Context, string, string) (string, error)
	SignUp(context.Context, model.User) error
	// Взаимодействие с публикациями
	GetPublication(context.Context, uint, *token.Claims) (model.Publication, error)
	CreatePublication(context.Context, model.Publication, *token.Claims) error
	GetAllPublications(context.Context) ([]model.Publication, error)
	UpdatePublication(context.Context, model.Publication, *token.Claims) error
	DeletePublication(context.Context, uint, *token.Claims) error
	// Взаимодействие с комментариями
	GetComment(context.Context, uint, *token.Claims) (model.Comment, error)
	GetCommentsFor(context.Context, uint) ([]model.Comment, error)
	CreateComment(context.Context, model.Comment, *token.Claims) error
	UpdateComment(context.Context, model.Comment, *token.Claims) error
	DeleteComment(context.Context, uint, *token.Claims) error
}
