package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type customerInteractor struct {
	CustomerPresenter  presenter.CustomerPresenter
	CustomerRepository repository.CustomerRepository
}

type CustomerInteractor interface {
	CreateCustomer(cus *entities.Customers) (*entities.Customers, error)
	ReadCustomer() ([]*entities.Customers, error)
	ReadCustomerById(ID string) (*entities.Customers, error)
	UpdateCustomer(cus *entities.Customers, ID string) (*entities.Customers, error)
	DeleteCustomer(ID string) error
}

func NewCustomerInteractor(cp presenter.CustomerPresenter, cr repository.CustomerRepository) CustomerInteractor {
	return &customerInteractor{
		CustomerPresenter:  cp,
		CustomerRepository: cr,
	}
}

func (ci *customerInteractor) CreateCustomer(cus *entities.Customers) (*entities.Customers, error) {
	createCus, err := ci.CustomerRepository.CreateCustomer(cus)

	if err != nil {
		return nil, err
	}

	return createCus, nil
}

func (ci *customerInteractor) ReadCustomer() ([]*entities.Customers, error) {
	readCus, err := ci.CustomerRepository.ReadCustomer()

	if err != nil {
		return nil, err
	}

	return readCus, nil

}

func (ci *customerInteractor) ReadCustomerById(ID string) (*entities.Customers, error) {

	readCusById, err := ci.CustomerRepository.ReadCustomerById(ID)

	if err != nil {
		return nil, err
	}

	return readCusById, nil
}

func (ci *customerInteractor) UpdateCustomer(cus *entities.Customers, ID string) (*entities.Customers, error) {

	updateCus, err := ci.CustomerRepository.UpdateCustomer(cus, ID)

	if err != nil {
		return nil, err
	}

	return updateCus, nil
}

func (ci *customerInteractor) DeleteCustomer(ID string) error {

	err := ci.CustomerRepository.DeleteCustomer(ID)

	if err != nil {
		return err
	}

	return nil
}
