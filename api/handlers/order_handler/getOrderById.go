package orderhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (od *orderHandler) GetOrderById(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		p := presenter.NewOrderPresenter()
		handle := handlers.NewAppHandler()

		if !handle.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.OrderErrorResponse(errors.New("invalid params id")))
		}

		order, err := ca.Order.ReadOrderById(id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.OrderErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.OrderSuccessResponse(order))
	}
}
