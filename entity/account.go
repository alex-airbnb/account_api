package entity

import (
	"errors"

	"github.com/go-playground/validator"
)

// ErrInvalidEmail Invalid Email Error
var ErrInvalidEmail = errors.New("Invalid Email")

// Account Entity
type Account struct {
	Email    string `validate:"email"`
	LastName string
	Name     string
}

// CreateAccount Create a new Account
func CreateAccount(email, lastName, name string) (Account, error) {
	validate := validator.New()
	account := Account{
		Email:    email,
		LastName: lastName,
		Name:     name,
	}

	if err := validate.Struct(account); err != nil {
		return Account{}, ErrInvalidEmail
	}

	return account, nil
}
