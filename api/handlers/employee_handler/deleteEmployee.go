package employeehandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *employeeHandlers) DeleteEmployee(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		p := presenter.NewEmployeePresenter()

		if !h.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.EmployeeErrorResponse(errors.New("invalid params id")))
		}

		if err := ca.Employee.DeleteEmployee(id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.EmployeeErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.EmployeeDeleteResponse())
	}
}
