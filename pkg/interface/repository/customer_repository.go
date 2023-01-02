package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &customerRepository{db}
}

func (cr *customerRepository) ReadCustomer() ([]*entities.Customers, error) {
	var customers []*entities.Customers

	if err := cr.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (cr *customerRepository) ReadCustomerById(ID string) (*entities.Customers, error) {
	var customer *entities.Customers

	if err := cr.db.First(&customer, "customer_id = ?", ID).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (cr *customerRepository) CreateCustomer(cus *entities.Customers) (*entities.Customers, error) {

	cus.Customer_id = uuid.New()

	if err := cr.db.Create(cus).Error; err != nil {
		return nil, err
	}

	return cus, nil
}

func (cr *customerRepository) UpdateCustomer(cus *entities.Customers, ID string) (*entities.Customers, error) {
	var oldCustomer *entities.Customers

	if err := cr.db.First(&oldCustomer, "customer_id = ?", ID).Error; err != nil {
		return nil, err
	}

	if oldCustomer == nil {
		return nil, errors.New("customer not found")
	}

	cus.Customer_id = oldCustomer.Customer_id

	if err := cr.db.Save(cus).Error; err != nil {
		return nil, err
	}

	return cus, nil
}

func (cr *customerRepository) DeleteCustomer(ID string) error {

	if err := cr.db.Where("customer_id = ?", ID).Delete(&entities.Customers{}).Error; err != nil {
		return err
	}

	return nil
}
