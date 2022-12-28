package employeehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *employeeHandlers) CreateEmployee(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var emp entities.Employees
		p := presenter.NewEmployeePresenter()

		if err := c.BodyParser(&emp); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.EmployeeErrorResponse(err))
		}

		createdEmployee, err := ca.Employee.CreateEmployee(&emp)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.EmployeeErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.EmployeeSuccessResponse(createdEmployee))
	}
}
