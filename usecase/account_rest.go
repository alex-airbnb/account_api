package usecase

import (
	"encoding/json"
	"errors"

	"github.com/alex-airbnb/account_api/adapter"
	"github.com/alex-airbnb/account_api/entity"
	"github.com/go-playground/validator"
)

var (
	// ErrInvalidJSONFormat Invalid JSON Format Error
	ErrInvalidJSONFormat = errors.New("Invalid JSON Format")

	// ErrMissingRequiredProperty Missing Required Property
	ErrMissingRequiredProperty = errors.New("Missing Required Property")
)

type accountREST struct {
	Repository adapter.RepositoryPort
}

// CreateAccountRequest Request for the CreateAccount method.
type CreateAccountRequest struct {
	Email    string `json:"email" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

// CreateAccountResponse Response for the CreateAccount method.
type CreateAccountResponse struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
}

// NewAccountREST Create a new instance of AccountREST Use Case.
func NewAccountREST(r adapter.RepositoryPort) UseCase {
	return &accountREST{
		Repository: r,
	}
}

// CreateAccount AccountREST Use Case method.
func (a *accountREST) CreateAccount(req []byte) (interface{}, error) {
	var createAccountRequest CreateAccountRequest
	validate := validator.New()

	if err := json.Unmarshal(req, &createAccountRequest); err != nil {
		return CreateAccountResponse{}, ErrInvalidJSONFormat
	}

	if err := validate.Struct(createAccountRequest); err != nil {
		return CreateAccountResponse{}, ErrMissingRequiredProperty
	}

	account, err := entity.CreateAccount(
		createAccountRequest.Email,
		createAccountRequest.LastName,
		createAccountRequest.Name,
	)

	if err != nil {
		return CreateAccountResponse{}, err
	}

	if err := a.Repository.Create(&account); err != nil {
		return CreateAccountResponse{}, err
	}

	return CreateAccountResponse{Email: account.Email, ID: account.ID}, nil
}
