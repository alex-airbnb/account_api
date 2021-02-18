package usecase

import (
	"testing"

	"github.com/alex-airbnb/account_api/entity"
	"github.com/franela/goblin"
)

func TestAccount(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("CreateAccountREST Use Case", func() {
		g.Describe("when the request is valid", func() {
			g.It("it should create an account and return a result", func() {
				a := &AccountREST{}
				expectedResult := CreateAccountResponse{
					Email: "account@mail.com",
					ID:    0,
				}
				req := []byte(`{
					"email": "account@mail.com",
					"lastName": "lastName",
					"name": "name"
				}`)

				currentResult, err := a.CreateAccount(req)

				g.Assert(currentResult).Equal(expectedResult)
				g.Assert(err).Equal(nil)
			})
		})

		g.Describe("when the request JSON is invalid", func() {
			g.It("it should return InvalidJSONFormat error", func() {
				a := &AccountREST{}
				req := []byte(`{
					"email": "account@mail.com",
					"lastName": "lastName",
					"name": "name",
				}`)

				currentResult, err := a.CreateAccount(req)

				g.Assert(currentResult).Equal(CreateAccountResponse{})
				g.Assert(err).Equal(ErrInvalidJSONFormat)
			})
		})

		g.Describe("when the request has a missing required property", func() {
			g.It("it should return MissingRequiredProperty error", func() {
				a := &AccountREST{}
				req := []byte(`{
					"email": "account@mail.com",
					"lastName": "lastName"
				}`)

				currentResult, err := a.CreateAccount(req)

				g.Assert(currentResult).Equal(CreateAccountResponse{})
				g.Assert(err).Equal(ErrMissingRequiredProperty)
			})
		})

		g.Describe("when the email is invalid", func() {
			g.It("it should return MissingRequiredProperty error", func() {
				a := &AccountREST{}
				req := []byte(`{
					"email": "account_mail.com",
					"lastName": "lastName",
					"name": "name"
				}`)

				currentResult, err := a.CreateAccount(req)

				g.Assert(currentResult).Equal(CreateAccountResponse{})
				g.Assert(err).Equal(entity.ErrInvalidEmail)
			})
		})
	})
}
