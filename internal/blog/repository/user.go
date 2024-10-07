package repository

import (
	"context"
	"tatKOM/model"
)

func (r *Repository) GetUser(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	result := r.DB.First(&user, id)
	return user, result.Error
}
func (r *Repository) CreateUser(ctx context.Context, user model.User) error {
	result := r.DB.Create(&user)
	return result.Error
}
func (r *Repository) UpdateUser(ctx context.Context, user model.User) error {
	result := r.DB.Save(&user)
	return result.Error
}
func (r *Repository) DeleteUser(ctx context.Context, id uint) error {
	var user model.User
	result := r.DB.Delete(&user, id)
	return result.Error
}
func (r *Repository) GetUserByLogin(ctx context.Context, login string) (model.User, error) {
	var user model.User
	result := r.DB.Where("login = ?", login).First(&user)
	return user, result.Error
}
