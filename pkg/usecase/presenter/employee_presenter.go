package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type EmployeePresenter interface {
	EmployeeSuccessResponse(emp *entities.Employees) *fiber.Map
	EmployeesSuccessResponse(emp []*entities.Employees) *fiber.Map
	EmployeeDeleteResponse() *fiber.Map
	EmployeeErrorResponse(err error) *fiber.Map
}
