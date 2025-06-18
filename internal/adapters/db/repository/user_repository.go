package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/adapters/db/sqlc"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/domain/entity"
	"github.com/pkg/errors"
)

type UserRepository struct {
	db      *pgxpool.Pool
	querier sqlc.Querier
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		querier: sqlc.New(db),
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	model, err := ur.querier.CreateUser(ctx, sqlc.CreateUserParams{
		Name:   user.Name,
		Email:  user.Email,
		Wallet: user.Wallet,
	})
	if err != nil {
		return nil, errors.Wrap(err, "db")
	}

	return &entity.User{
		ID:     model.ID,
		Name:   model.Name,
		Email:  model.Email,
		Wallet: model.Wallet,
	}, nil
}
func (ur *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := ur.querier.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "db")
	}

	return &entity.User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Wallet: user.Wallet,
	}, nil
}
func (ur *UserRepository) Update(ctx context.Context, user *entity.User) error {
	if err := ur.querier.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Wallet: user.Wallet,
	}); err != nil {
		return errors.Wrap(err, "db")
	}

	return nil
}
func (ur *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := ur.querier.DeleteUser(ctx, id); err != nil {
		return errors.Wrap(err, "db")
	}

	return nil
}
