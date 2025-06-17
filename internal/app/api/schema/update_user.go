package schema

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UpdateUserRequest struct {
	ID     string          `param:"id"`
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Wallet decimal.Decimal `json:"wallet"`
}

func (r UpdateUserRequest) Validate() error {
	_, err := uuid.Parse(r.ID)
	if err != nil {
		return NewValidationError("id is required").WithField("id")
	}

	if r.Name == "" {
		return NewValidationError("name is required").WithField("name")
	}

	if r.Email == "" {
		return NewValidationError("email is required").WithField("email")
	}

	if r.Wallet.IsZero() {
		return NewValidationError("wallet must be greater than zero").WithField("wallet")
	}

	return nil
}

type UpdateUserResponse struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Wallet decimal.Decimal `json:"wallet"`
}
