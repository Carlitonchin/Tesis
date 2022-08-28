package repository

import (
	"context"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (s *UserRepository) Create(ctx context.Context, user *model.User) error {
	result := s.DB.Create(user)

	return result.Error
}

func (s *UserRepository) FindById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	result := s.DB.First(&user, id)

	return &user, result.Error
}