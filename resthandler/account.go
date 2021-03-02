package resthandler

import (
	"net/http"

	"github.com/alex-airbnb/account_api/usecase"
	"github.com/gofiber/fiber/v2"
)

// AccountRouter Set up all the handlers for /account route.
func AccountRouter(r fiber.Router, a usecase.UseCase) {
	r.Post("/", createAccount(a))
}

func createAccount(a usecase.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusCreated).JSON(response{
			Data:       0,
			Error:      false,
			Message:    "Account Created Successfully",
			Status:     "Created",
			StatusCode: http.StatusCreated,
		})
	}
}
