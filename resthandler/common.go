package resthandler

import (
	"github.com/alex-airbnb/account_api/usecase"
	"github.com/gofiber/fiber/v2"
)

var createMockFunction func(model interface{}) error

type repositoryMock struct{}

func (r *repositoryMock) Create(model interface{}) error {
	return createMockFunction(model)
}

type response struct {
	Data       int    `json:"data"`
	Error      bool   `json:"error"`
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
}

// Setup Set up a new Fiber App.
func Setup() *fiber.App {
	app := fiber.New()
	accountUseCase := usecase.NewAccountREST(&repositoryMock{})
	accountGroup := app.Group("/account")

	AccountRouter(accountGroup, accountUseCase)

	return app
}
