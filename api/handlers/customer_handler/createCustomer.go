package customerhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *customerHandlers) CreateCustomer(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cus *entities.Customers
		p := presenter.NewCustomerPresenter()

		if err := c.BodyParser(&cus); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.CustomerErrorResponse(err))
		}

		customer, err := ca.Customer.CreateCustomer(cus)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.CustomerErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.CustomerSuccessResponse(customer))
	}
}
