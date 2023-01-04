package employeehandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	"gorm.io/gorm"
)

func (h *employeeHandlers) UpdateEmployee(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		p := presenter.NewEmployeePresenter()
		handle := handlers.NewAppHandler()
		var emp entities.Employees

		if !handle.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.EmployeeErrorResponse(errors.New("invalid params id")))
		}

		if err := c.BodyParser(&emp); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.EmployeeErrorResponse(err))
		}

		_, err := ca.Employee.UpdateEmployee(&emp, id)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(p.EmployeeErrorResponse(errors.New("record not found")))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(p.EmployeeErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.EmployeeUpdateResponse())

	}
}
