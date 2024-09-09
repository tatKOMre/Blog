package repository

import (
	"context"
	"tatKOM/model"
)

func (r *Repository) GetPublication(ctx context.Context, id uint) (model.Publication, error) {
	var publication model.Publication
	result := r.DB.First(&publication, id)
	return publication, result.Error
}
func (r *Repository) CreatePublication(ctx context.Context, publication model.Publication) error {
	result := r.DB.Create(&publication)
	return result.Error
}
func (r *Repository) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	var publications []model.Publication
	result := r.DB.Find(&publications)
	return publications, result.Error
}
func (r *Repository) UpdatePublication(ctx context.Context, publication model.Publication) error {
	result := r.DB.Save(&publication)
	return result.Error
}
func (r *Repository) DeletePublication(ctx context.Context, id uint) error {
	var publication model.Publication
	result := r.DB.Delete(&publication, id)
	return result.Error
}
