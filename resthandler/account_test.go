package resthandler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gofiber/fiber/v2"
)

func TestAccountRESTHandler(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Account REST Handler", func() {
		g.Describe("POST /account", func() {
			g.Describe("when the body is valid", func() {
				g.It("it should response Created with the account info", func() {
					var restHandler *fiber.App = Setup()
					var currentResponseBody response
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

					res, err := restHandler.Test(req, -1)
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
