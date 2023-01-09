package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type Order struct {
	Order_id    uuid.UUID `json:"order_id"`
	Customer_id uuid.UUID `json:"customer_id"`
	Order_date  string    `json:"order_date"`

	OrderDetails []entities.OrderDetails `gorm:"foreignkey:Order_id"`
}

type orderPresenter struct {
}

func NewOrderPresenter() presenter.OrderPresenter {
	return &orderPresenter{}
}

func (op *orderPresenter) OrderSuccessResponse(order *entities.Orders) *fiber.Map {
	var o = &Order{
		Order_id:     order.Order_id,
		Customer_id:  order.Customer_id,
		OrderDetails: order.OrderDetails,
	}

	return &fiber.Map{
		"status": true,
		"data":   o,
		"error":  nil,
	}
}

func (op *orderPresenter) OrderErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}

func (op *orderPresenter) OrderDeleteResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   nil,
		"error":  nil,
	}
}
