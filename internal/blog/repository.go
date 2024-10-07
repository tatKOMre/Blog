package blog

import (
	"context"

	"tatKOM/model"
)

type Repository interface {
	// Взаимодействие с юзером
	GetUser(context.Context, uint) (model.User, error)
	CreateUser(context.Context, model.User) error
	UpdateUser(context.Context, model.User) error
	DeleteUser(context.Context, uint) error
	GetUserByLogin(context.Context, string) (model.User, error)
	// Взаимодействие с публикациями
	GetPublication(context.Context, uint) (model.Publication, error)
	GetAllPublications(context.Context) ([]model.Publication, error)
	CreatePublication(context.Context, model.Publication) error
	UpdatePublication(context.Context, model.Publication) error
	DeletePublication(context.Context, uint) error
	// Взаимодействие с комментариями
	GetComment(context.Context, uint) (model.Comment, error)
	GetCommentsFor(context.Context, uint) ([]model.Comment, error)
	CreateComment(context.Context, model.Comment) error
	UpdateComment(context.Context, model.Comment) error
	DeleteComment(context.Context, uint) error
}
