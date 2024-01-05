package domain

import (
	"context"
	"prototype/domain/user/models"
)

// interface for repository
type IUserMysqlRepository interface {
	Fetch(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user models.User) (models.User, error)
	Update(ctx context.Context, user models.User) (models.User, error)
	GetByID(ctx context.Context, id uint) (models.User, error)
	Delete(ctx context.Context, id uint) error
}

// interface for usecase
type IUserUsecase interface {
	Fetch(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user models.User) (models.User, error)
	Update(ctx context.Context, user models.User) (models.User, error)
	GetByID(ctx context.Context, id uint) (models.User, error)
	Delete(ctx context.Context, id uint) error
}
