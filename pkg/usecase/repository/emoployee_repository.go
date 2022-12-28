package repository

import "github.com/nutthanonn/go-clean-architecture/pkg/entities"

type EmployeeRepository interface {
	CreateEmployee(emp *entities.Employees) (*entities.Employees, error)
	ReadEmployee() ([]*entities.Employees, error)
	ReadEmployeeById(ID string) (*entities.Employees, error)
	UpdateEmployee(emp *entities.Employees, ID string) (*entities.Employees, error)
	DeleteEmployee(ID string) error
}
