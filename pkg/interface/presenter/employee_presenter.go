package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type Employee struct {
	Employee_id uuid.UUID `json:"employee_id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Email       string    `json:"email"`
	Job_title   string    `json:"job_title"`
}

type employeePresenter struct{}

func NewEmployeePresenter() presenter.EmployeePresenter {
	return &employeePresenter{}
}

func (ep *employeePresenter) EmployeeSuccessResponse(emp *entities.Employees) *fiber.Map {
	var res = Employee{
		Employee_id: emp.Employee_id,
		First_name:  emp.First_name,
		Last_name:   emp.Last_name,
		Email:       emp.Email,
		Job_title:   emp.Job_title,
	}

	return &fiber.Map{
		"data":   res,
		"error":  nil,
		"status": true,
	}
}

func (ep *employeePresenter) EmployeesSuccessResponse(emp []*entities.Employees) *fiber.Map {
	var res []Employee

	for _, v := range emp {
		emp := Employee{
			Employee_id: v.Employee_id,
			First_name:  v.First_name,
			Last_name:   v.Last_name,
			Email:       v.Email,
			Job_title:   v.Job_title,
		}
		res = append(res, emp)
	}

	return &fiber.Map{
		"data":   res,
		"error":  nil,
		"status": true,
	}
}

func (ep *employeePresenter) EmployeeUpdateResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   nil,
		"error":  nil,
	}
}

func (ep *employeePresenter) EmployeeDeleteResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   nil,
		"error":  nil,
	}
}

func (ep *employeePresenter) EmployeeErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
