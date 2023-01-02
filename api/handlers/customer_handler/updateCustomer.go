package customerhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *customerHandlers) UpdateCustomer(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cus *entities.Customers
		id := c.Params("id")
		p := presenter.NewCustomerPresenter()
		handle := handlers.NewAppHandler()

		if !handle.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.CustomerErrorResponse(errors.New("invalid params id")))
		}

		if err := c.BodyParser(&cus); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.CustomerErrorResponse(err))
		}

		customer, err := ca.Customer.UpdateCustomer(cus, id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.CustomerErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.CustomerSuccessResponse(customer))
	}
}
