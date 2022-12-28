package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) repository.EmployeeRepository {
	return &employeeRepository{db}
}

func (er *employeeRepository) CreateEmployee(emp *entities.Employees) (*entities.Employees, error) {
	emp.Employee_id = uuid.New()
	emp.Create_at = time.Now()
	emp.Update_at = time.Now()
	emp.Hire_date = time.Now()

	if err := er.db.Create(&emp).Error; err != nil {
		return nil, err
	}

	return emp, nil
}

func (er *employeeRepository) ReadEmployee() ([]*entities.Employees, error) {
	var emps []*entities.Employees

	if err := er.db.Find(&emps).Error; err != nil {
		return nil, err
	}

	return emps, nil
}

func (er *employeeRepository) ReadEmployeeById(ID string) (*entities.Employees, error) {
	var emp *entities.Employees

	if err := er.db.First(&emp, "employee_id = ?", ID).Error; err != nil {
		return nil, err
	}

	return emp, nil
}

func (er *employeeRepository) UpdateEmployee(emp *entities.Employees, ID string) (*entities.Employees, error) {
	var oldEmp entities.Employees
	if err := er.db.First(&oldEmp, "employee_id = ?", ID).Error; err != nil {
		return nil, err
	}

	emp.Employee_id = oldEmp.Employee_id
	emp.Create_at = oldEmp.Create_at
	emp.Update_at = oldEmp.Update_at

	if err := er.db.Save(&emp).Error; err != nil {
		return nil, err
	}

	return emp, nil
}

func (er *employeeRepository) DeleteEmployee(ID string) error {
	if err := er.db.Where("employee_id = ?", ID).Delete(&entities.Employees{}).Error; err != nil {
		return err
	}

	return nil
}
