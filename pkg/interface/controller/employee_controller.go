package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type employeeController struct {
	EmployeeInteractor interactor.EmployeeInteractor
}

type EmployeeController interface {
	CreateEmployee(emp *entities.Employees) (*entities.Employees, error)
	ReadEmployee() ([]*entities.Employees, error)
	ReadEmployeeById(ID string) (*entities.Employees, error)
	UpdateEmployee(emp *entities.Employees, ID string) (*entities.Employees, error)
	DeleteEmployee(ID string) error
}

func NewEmployeeController(ei interactor.EmployeeInteractor) EmployeeController {
	return &employeeController{ei}
}

func (ec *employeeController) CreateEmployee(emp *entities.Employees) (*entities.Employees, error) {
	return ec.EmployeeInteractor.CreateEmployee(emp)
}

func (ec *employeeController) ReadEmployee() ([]*entities.Employees, error) {
	return ec.EmployeeInteractor.ReadEmployee()
}

func (ec *employeeController) ReadEmployeeById(ID string) (*entities.Employees, error) {
	return ec.EmployeeInteractor.ReadEmployeeById(ID)
}

func (ec *employeeController) UpdateEmployee(emp *entities.Employees, ID string) (*entities.Employees, error) {
	return ec.EmployeeInteractor.UpdateEmployee(emp, ID)
}

func (ec *employeeController) DeleteEmployee(ID string) error {
	return ec.EmployeeInteractor.DeleteEmployee(ID)
}
