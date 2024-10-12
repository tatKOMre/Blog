package service

import (
	"context"
	"log"
	"tatKOM/model"
	"tatKOM/pkg/hash"
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
	user.Password = hash.Hash(user.Password)
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
func (s *Service) Singin(ctx context.Context, login string, password string) (string, error) {
	user, err := s.Repository.GetUserByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	if hash.Hash(password) != user.Password {
		return "", wrongpass
	}
	claims := token.Claims{
		ID:         user.ID,
		Name:       user.Name,
		Password:   user.Password,
		Permission: user.Permission,
	}

	tkn, err := token.GenerateJWT(&claims, s.SignKey)
	if err != nil {
		return "", err
	}
	return tkn, err
}
