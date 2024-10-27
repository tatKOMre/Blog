package service

import (
	"context"
	"log"
	"tatKOM/model"
	"tatKOM/pkg/token"
)

func (s *Service) GetPublication(ctx context.Context, id uint, act *token.Claims) (model.Publication, error) {
	user, err := s.Repository.GetPublication(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if !act.Permission {
		return model.Publication{}, errNotPermission
	}
	return user, nil
}
func (s *Service) CreatePublication(ctx context.Context, publication model.Publication, act *token.Claims) error {

	if !act.Permission {
		return errNotPermission
	}
	result := s.Repository.CreatePublication(ctx, publication)
	return result
}
func (s *Service) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	return s.Repository.GetAllPublications(ctx)
}
func (s *Service) UpdatePublication(ctx context.Context, publication model.Publication, act *token.Claims) error {
	if !act.Permission {
		return errNotPermission
	}
	result := s.Repository.UpdatePublication(ctx, publication)
	return result
}
func (s *Service) DeletePublication(ctx context.Context, id uint, act *token.Claims) error {
	if !act.Permission {
		return errNotPermission
	}
	result := s.Repository.DeletePublication(ctx, id)
	return result
}
