package repository

import "github.com/nutthanonn/go-clean-architecture/pkg/entities"

type CustomerRepository interface {
	ReadCustomer() ([]*entities.Customers, error)
	ReadCustomerById(ID string) (*entities.Customers, error)
	CreateCustomer(cus *entities.Customers) (*entities.Customers, error)
	UpdateCustomer(cus *entities.Customers, ID string) (*entities.Customers, error)
	DeleteCustomer(ID string) error
}
