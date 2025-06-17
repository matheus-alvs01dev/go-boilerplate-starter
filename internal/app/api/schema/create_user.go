package schema

import "github.com/shopspring/decimal"

type CreateUserRequest struct {
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Wallet decimal.Decimal `json:"wallet"`
}

func (r CreateUserRequest) Validate() error {
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

type CreateUserResponse struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Wallet decimal.Decimal `json:"wallet"`
}
