package orderhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (oh *orderHandler) CreateOrder(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var order entities.Orders

		fields := []string{"OrderDetails", "Customer_id"}

		p := presenter.NewOrderPresenter()
		handle := handlers.NewAppHandler()

		if err := c.BodyParser(&order); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.OrderErrorResponse(err))
		}

		if err := handle.ValidateFields(fields, order); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.OrderErrorResponse(err))
		}

		createdOrder, err := ca.Order.CreateOrder(&order)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.OrderErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.OrderSuccessResponse(createdOrder))

	}
}
