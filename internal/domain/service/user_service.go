package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/domain/entity"
	"github.com/pkg/errors"
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
	user, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "repo")
	}

	return user, nil
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "repo")
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, user *entity.User) error {
	if err := s.repo.Update(ctx, user); err != nil {
		return errors.Wrap(err, "repo")
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "repo")
	}

	return nil
}
