package employeehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *employeeHandlers) GetEmployee(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := presenter.NewEmployeePresenter()
		var emps []*entities.Employees

		emps, err := ca.Employee.ReadEmployee()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.EmployeeErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.EmployeesSuccessResponse(emps))
	}
}
