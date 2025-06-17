package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Wallet    decimal.Decimal
}

func NewUser(name, email string, wallet decimal.Decimal) *User {
	return &User{
		Name:   name,
		Email:  email,
		Wallet: wallet,
	}
}
