package customerhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	"gorm.io/gorm"
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

		_, err := ca.Customer.UpdateCustomer(cus, id)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(p.CustomerErrorResponse(errors.New("record not found")))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(p.CustomerErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.CustomerUpdateResponse())
	}
}
