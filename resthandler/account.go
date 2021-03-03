package resthandler

import (
	"log"
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
		account, err := a.CreateAccount(c.Body())

		if err != nil && err.Error() == "Repository Create Error" {
			log.Print(err)

			return c.Status(http.StatusInternalServerError).JSON(response{
				Data:       account,
				Error:      true,
				Message:    "Server Error",
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
		}

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(response{
				Data:       account,
				Error:      true,
				Message:    err.Error(),
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
			})
		}

		return c.Status(http.StatusCreated).JSON(response{
			Data:       account,
			Error:      false,
			Message:    "Account Created Successfully",
			Status:     "Created",
			StatusCode: http.StatusCreated,
		})
	}
}
