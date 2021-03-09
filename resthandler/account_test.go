package resthandler

import (
	"bytes"
	"encoding/json"
	"errors"
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
		g.Describe("POST /api/v1/account", func() {
			var currentResponseBody response
			var a usecase.UseCase = usecase.NewAccountREST(&repositoryMock{})
			createMockFunction = func(m interface{}) error {
				return nil
			}
			var r *fiber.App = NewRESTHandler(a)

			g.Describe("when the request is valid", func() {
				g.It("it should response Created with the account info", func() {
					expectedResponseBody := response{
						Data: map[string]interface{}{
							"email": "account@mail.com",
							"id":    0,
						},
						Error:      false,
						Message:    "Account Created Successfully",
						Status:     "Created",
						StatusCode: http.StatusCreated,
					}
					req, _ := http.NewRequest(
						"POST",
						"/api/v1/account",
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

			g.Describe("when the request is invalid", func() {
				g.It("it should response Bad Request", func() {
					expectedResponseBody := response{
						Data: map[string]interface{}{
							"email": "",
							"id":    0,
						},
						Error:      true,
						Message:    "Invalid JSON Format",
						Status:     "Bad Request",
						StatusCode: http.StatusBadRequest,
					}
					req, _ := http.NewRequest(
						"POST",
						"/api/v1/account",
						bytes.NewBuffer([]byte(`{
							"email": "account@mail.com",
							"lastName": "lastName",
							"name": "name",
						}`)),
					)

					res, err := r.Test(req, -1)
					body, _ := ioutil.ReadAll(res.Body)

					json.Unmarshal(body, &currentResponseBody)

					g.Assert(res.StatusCode).Equal(http.StatusBadRequest)
					g.Assert(currentResponseBody).Equal(expectedResponseBody)
					g.Assert(err).Equal(nil)
				})
			})

			g.Describe("when the request has missed a parameter", func() {
				g.It("it should response Bad Request", func() {
					expectedResponseBody := response{
						Data: map[string]interface{}{
							"email": "",
							"id":    0,
						},
						Error:      true,
						Message:    "Missing Required Property",
						Status:     "Bad Request",
						StatusCode: http.StatusBadRequest,
					}
					req, _ := http.NewRequest(
						"POST",
						"/api/v1/account",
						bytes.NewBuffer([]byte(`{
							"email": "account@mail.com",
							"lastName": "lastName",
							"nam": "name"
						}`)),
					)

					res, err := r.Test(req, -1)
					body, _ := ioutil.ReadAll(res.Body)

					json.Unmarshal(body, &currentResponseBody)

					g.Assert(res.StatusCode).Equal(http.StatusBadRequest)
					g.Assert(currentResponseBody).Equal(expectedResponseBody)
					g.Assert(err).Equal(nil)
				})
			})

			g.Describe("when the repository fails", func() {
				g.It("it should response Internal Server Error", func() {
					createMockFunction = func(m interface{}) error {
						return errors.New("Repository Create Error")
					}
					expectedResponseBody := response{
						Data: map[string]interface{}{
							"email": "",
							"id":    0,
						},
						Error:      true,
						Message:    "Server Error",
						Status:     "Internal Server Error",
						StatusCode: http.StatusInternalServerError,
					}
					req, _ := http.NewRequest(
						"POST",
						"/api/v1/account",
						bytes.NewBuffer([]byte(`{
							"email": "account@mail.com",
							"lastName": "lastName",
							"name": "name"
						}`)),
					)

					res, err := r.Test(req, -1)
					body, _ := ioutil.ReadAll(res.Body)

					json.Unmarshal(body, &currentResponseBody)

					g.Assert(res.StatusCode).Equal(http.StatusInternalServerError)
					g.Assert(currentResponseBody).Equal(expectedResponseBody)
					g.Assert(err).Equal(nil)
				})
			})
		})
	})
}
