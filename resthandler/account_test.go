package resthandler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alex-airbnb/account_api/usecase"
	"github.com/franela/goblin"
	"github.com/gofiber/fiber/v2"
)

var createMockFunction func(model interface{}) error

type repositoryMock struct{}

func (r *repositoryMock) Create(model interface{}) error {
	return createMockFunction(model)
}

func TestAccountRESTHandler(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Account REST Handler", func() {
		g.Describe("POST /account", func() {
			g.Describe("when the body is valid", func() {
				g.It("it should response Created with the account info", func() {
					var currentResponseBody response
					var a usecase.UseCase = usecase.NewAccountREST(&repositoryMock{})
					createMockFunction = func(m interface{}) error {
						return nil
					}
					var r *fiber.App = NewRESTHandler(a)
					expectedResponseBody := response{
						Data:       0,
						Error:      false,
						Message:    "Account Created Successfully",
						Status:     "Created",
						StatusCode: http.StatusCreated,
					}
					req, _ := http.NewRequest(
						"POST",
						"/account",
						bytes.NewBuffer([]byte(`{
							"email": "account@mail.com",
							"lastName": "lastName",
							"name": "name"
						}`)),
					)

					res, err := r.Test(req, -1)
					body, _ := ioutil.ReadAll(res.Body)

					json.Unmarshal(body, &currentResponseBody)

					g.Assert(res.StatusCode).Equal(http.StatusCreated)
					g.Assert(currentResponseBody).Equal(expectedResponseBody)
					g.Assert(err).Equal(nil)
				})
			})
		})
	})
}
