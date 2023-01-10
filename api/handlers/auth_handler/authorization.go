package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func (h *authHandler) Authorization(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Authorization")
	}
}
