package authhandler

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/models"
)

func (h *authHandler) SignIn(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.SignIn
		p := presenter.NewAuthPresenter()
		handle := handlers.NewAppHandler()
		fields := []string{"Email", "Password"}

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.AuthErrorResponse(err))
		}

		if err := handle.ValidateFields(fields, user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.AuthErrorResponse(err))
		}

		ok, err := ca.Auth.SignIn(&user)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.AuthErrorResponse(err))
		}

		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(p.AuthErrorResponse(errors.New("invalid email or password")))
		}

		token, err := ca.Auth.CreateToken(&entities.Customers{Email: user.Email})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.AuthErrorResponse(err))
		}

		c.Cookie(&fiber.Cookie{
			Name:     "access_token",
			Value:    token.AccessToken,
			HTTPOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		})

		return c.Status(fiber.StatusOK).JSON(p.AuthResponse())
	}
}
