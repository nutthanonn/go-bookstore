package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type employeeInteractor struct {
	EmployeePresenter  presenter.EmployeePresenter
	EmployeeRepository repository.EmployeeRepository
}

type EmployeeInteractor interface {
	CreateEmployee(emp *entities.Employees) (*entities.Employees, error)
	ReadEmployee() ([]*entities.Employees, error)
	UpdateEmployee(emp *entities.Employees, ID string) (*entities.Employees, error)
	DeleteEmployee(ID string) error
	ReadEmployeeById(ID string) (*entities.Employees, error)
}

func NewEmployeeInteractor(up presenter.EmployeePresenter, ur repository.EmployeeRepository) EmployeeInteractor {
	return &employeeInteractor{
		EmployeeRepository: ur,
		EmployeePresenter:  up,
	}
}

func (ei *employeeInteractor) CreateEmployee(emp *entities.Employees) (*entities.Employees, error) {
	createdEmp, err := ei.EmployeeRepository.CreateEmployee(emp)
	if err != nil {
		return nil, err
	}

	return createdEmp, nil
}

func (ei *employeeInteractor) ReadEmployee() ([]*entities.Employees, error) {
	emps, err := ei.EmployeeRepository.ReadEmployee()

	if err != nil {
		return nil, err
	}

	return emps, nil
}

func (ei *employeeInteractor) ReadEmployeeById(ID string) (*entities.Employees, error) {
	emp, err := ei.EmployeeRepository.ReadEmployeeById(ID)

	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (ei *employeeInteractor) UpdateEmployee(emp *entities.Employees, ID string) (*entities.Employees, error) {
	emp, err := ei.EmployeeRepository.UpdateEmployee(emp, ID)

	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (ei *employeeInteractor) DeleteEmployee(ID string) error {
	if err := ei.EmployeeRepository.DeleteEmployee(ID); err != nil {
		return err
	}

	return nil
}
