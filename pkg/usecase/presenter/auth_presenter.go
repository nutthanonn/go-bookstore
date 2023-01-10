package presenter

import (
	"github.com/gofiber/fiber/v2"
)

type AuthPresenter interface {
	AuthResponse() *fiber.Map
	AuthErrorResponse(err error) *fiber.Map
}
