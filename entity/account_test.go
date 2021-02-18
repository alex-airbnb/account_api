package entity

import (
	"testing"

	"github.com/franela/goblin"
)

func TestAccount(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("CreateAccount Entity", func() {
		g.Describe("when the email, lastName, and name are valid", func() {
			g.It("it should create a account", func() {
				expectedAccount := Account{
					Email:    "account@mail.com",
					LastName: "Account Last Name",
					Name:     "Account Name",
				}

				currentAccount, err := CreateAccount(
					expectedAccount.Email,
					expectedAccount.LastName,
					expectedAccount.Name,
				)

				g.Assert(currentAccount).Equal(expectedAccount)
				g.Assert(err).Equal(nil)
			})
		})

		g.Describe("when the email is invalid", func() {
			g.It("it should return InvalidEmail error", func() {
				currentAccount, err := CreateAccount(
					"Email",
					"Last Name",
					"Name",
				)

				g.Assert(currentAccount).Equal(Account{})
				g.Assert(err).Equal(ErrInvalidEmail)
			})
		})
	})
}
