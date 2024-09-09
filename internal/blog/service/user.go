package service

import (
	"context"
	"log"
	"tatKOM/model"
	"tatKOM/pkg/token"
)

func (s *Service) GetUser(ctx context.Context, id uint, act *token.Claims) (model.User, error) {
	user, err := s.Repository.GetUser(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if !act.Permission {
		return model.User{}, errNotPermission
	}
	return user, nil
}
func (s *Service) CreateUser(ctx context.Context, user model.User, act *token.Claims) error {

	if !act.Permission {
		return errNotPermission
	}
	result := s.Repository.CreateUser(ctx, user)
	return result
}
func (s *Service) UpdateUser(ctx context.Context, user model.User, act *token.Claims) error {
	if !act.Permission {
		return errNotPermission
	}
	result := s.Repository.UpdateUser(ctx, user)
	return result
}
func (s *Service) DeleteUser(ctx context.Context, id uint, act *token.Claims) error {
	if !act.Permission {
		return errNotPermission
	}
	result := s.Repository.DeleteUser(ctx, id)
	return result
}
