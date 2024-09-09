package repository

import (
	"context"
	"tatKOM/model"
)

func (r *Repository) GetComment(ctx context.Context, id uint) (model.Comment, error) {
	var comment model.Comment
	result := r.DB.First(&comment, id)
	return comment, result.Error
}
func (r *Repository) GetCommentsFor(ctx context.Context, id uint) ([]model.Comment, error) {
	var comments []model.Comment
	result := r.DB.Where("publication_id = ?", id).Find(&comments)
	return comments, result.Error
}
func (r *Repository) CreateComment(ctx context.Context, comment model.Comment) error {
	result := r.DB.Create(&comment)
	return result.Error
}
func (r *Repository) UpdateComment(ctx context.Context, comment model.Comment) error {
	result := r.DB.Save(&comment)
	return result.Error
}
func (r *Repository) DeleteComment(ctx context.Context, id uint) error {
	var comment model.Comment
	result := r.DB.Delete(&comment, id)
	return result.Error
}
