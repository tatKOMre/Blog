package service

import (
	"context"
	"log"
	"tatKOM/model"
	"tatKOM/pkg/token"
)

func (s *Service) GetComment(ctx context.Context, id uint, act *token.Claims) (model.Comment, error) {
	comment, err := s.Repository.GetComment(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if !act.Permission {
		return model.Comment{}, errNotPermission
	}
	return comment, nil
}
func (s *Service) CreateComment(ctx context.Context, comment model.Comment, act *token.Claims) error {

	if !act.Permission {
		return errNotPermission
	}
	result := s.Repository.CreateComment(ctx, comment)
	return result
}
func (s *Service) GetCommentsFor(ctx context.Context, id uint) {
	s.Repository.GetCommentsFor(ctx, id)
}
func (s *Service) UpdateComment(ctx context.Context, comment model.Comment, act *token.Claims) error {
	if act.ID != comment.UserID && !act.Permission {
		return errNotPermission
	}
	result := s.Repository.UpdateComment(ctx, comment)
	return result
}
func (s *Service) DeleteComment(ctx context.Context, id uint, act *token.Claims) error {
	if act.ID != id && !act.Permission {
		return errNotPermission
	}
	result := s.Repository.DeleteComment(ctx, id)
	return result
}
