package customerhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *customerHandlers) DeleteCustomer(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := ca.Customer.DeleteCustomer(id)
		p := presenter.NewCustomerPresenter()
		if !h.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.CustomerErrorResponse(errors.New("invalid params id")))
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.CustomerErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.CustomerDeleteSuccessResponse())
	}
}