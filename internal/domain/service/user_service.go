package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	panic("implement me")
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	panic("implement me")
}

func (s *UserService) Update(ctx context.Context, user *entity.User) error {
	panic("implement me")
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
