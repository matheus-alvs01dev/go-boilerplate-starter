package schema

import (
	"github.com/matheus-alvs01dev/go-boilerplate/internal/domain/entity"
	"github.com/shopspring/decimal"
)

type CreateUserRequest struct {
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Wallet decimal.Decimal `json:"wallet"`
}

func (r CreateUserRequest) Validate() error {
	if r.Name == "" {
		return NewValidationError("name is required").WithField(r, "name")
	}

	if r.Email == "" {
		return NewValidationError("email is required").WithField(r, "email")
	}

	if r.Wallet.LessThanOrEqual(decimal.Zero) {
		return NewValidationError("wallet must be greater than zero").WithField(r, "wallet")
	}

	return nil
}

type CreateUserResponse struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Wallet decimal.Decimal `json:"wallet"`
}

func NewCreateUserResponse(user *entity.User) CreateUserResponse {
	return CreateUserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Wallet: user.Wallet,
	}
}
