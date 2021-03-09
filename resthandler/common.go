package resthandler

import (
	"github.com/alex-airbnb/account_api/usecase"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Data       interface{} `json:"data"`
	Error      bool        `json:"error"`
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
}

// NewRESTHandler Set up a new Fiber App.
func NewRESTHandler(accountUseCase usecase.UseCase) *fiber.App {
	app := fiber.New()
	accountGroup := app.Group("/api/v1/account")

	AccountRouter(accountGroup, accountUseCase)

	return app
}
