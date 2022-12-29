package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type Customer struct {
	Customer_id uuid.UUID `json:"customer_id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Zip         string    `json:"zip"`
}

type customerPresenter struct{}

func NewCustomerPresenter() presenter.CustomerPresenter {
	return &customerPresenter{}
}

func (cp *customerPresenter) CustomerSuccessResponse(cus *entities.Customers) *fiber.Map {
	var res = Customer{
		Customer_id: cus.Customer_id,
		First_name:  cus.First_name,
		Last_name:   cus.Last_name,
		Phone:       cus.Phone,
		Address:     cus.Address,
		City:        cus.City,
		State:       cus.State,
		Zip:         cus.Zip,
	}

	return &fiber.Map{
		"data":   res,
		"error":  nil,
		"status": true,
	}
}

func (cp *customerPresenter) CustomersSuccessResponse(cus []*entities.Customers) *fiber.Map {
	var res []Customer

	for _, v := range cus {
		cus := Customer{
			Customer_id: v.Customer_id,
			First_name:  v.First_name,
			Last_name:   v.Last_name,
			Phone:       v.Phone,
			Address:     v.Address,
			City:        v.City,
			State:       v.State,
			Zip:         v.Zip,
		}
		res = append(res, cus)
	}

	return &fiber.Map{
		"data":   res,
		"error":  nil,
		"status": true,
	}
}

func (cp *customerPresenter) CustomerErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":   nil,
		"error":  err.Error(),
		"status": false,
	}
}

func (cp *customerPresenter) CustomerDeleteSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"data":   nil,
		"error":  nil,
		"status": true,
	}
}
