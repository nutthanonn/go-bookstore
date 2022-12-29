package customerhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *customerHandlers) GetCustomer(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customers, err := ca.Customer.ReadCustomer()
		p := presenter.NewCustomerPresenter()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.CustomerErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.CustomersSuccessResponse(customers))
	}
}
