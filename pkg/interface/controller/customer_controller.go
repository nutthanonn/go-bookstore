package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type customerController struct {
	CustomerInteractor interactor.CustomerInteractor
}

type CustomerController interface {
	CreateCustomer(customer *entities.Customers) (*entities.Customers, error)
	ReadCustomer() ([]*entities.Customers, error)
	ReadCustomerById(ID string) (*entities.Customers, error)
	UpdateCustomer(customer *entities.Customers, ID string) (*entities.Customers, error)
	DeleteCustomer(ID string) error
}

func NewCustomerController(ci interactor.CustomerInteractor) CustomerController {
	return &customerController{ci}
}

func (cc *customerController) CreateCustomer(customer *entities.Customers) (*entities.Customers, error) {
	return cc.CustomerInteractor.CreateCustomer(customer)
}

func (cc *customerController) ReadCustomer() ([]*entities.Customers, error) {
	return cc.CustomerInteractor.ReadCustomer()
}

func (cc *customerController) ReadCustomerById(ID string) (*entities.Customers, error) {
	return cc.CustomerInteractor.ReadCustomerById(ID)
}

func (cc *customerController) UpdateCustomer(customer *entities.Customers, ID string) (*entities.Customers, error) {
	return cc.CustomerInteractor.UpdateCustomer(customer, ID)
}

func (cc *customerController) DeleteCustomer(ID string) error {
	return cc.CustomerInteractor.DeleteCustomer(ID)
}
