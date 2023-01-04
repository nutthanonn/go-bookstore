package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type orderController struct {
	OrderInteractor interactor.OrderInteractor
}

type OrderController interface {
	CreateOrder(order *entities.Orders) (*entities.Orders, error)
	ReadOrderById(order_id string) (*entities.Orders, error)
	UpdateOrder(order *entities.Orders, ID string) (*entities.Orders, error)
	DeleteOrder(order *entities.Orders, ID string) (*entities.Orders, error)
}

func NewOrderController(oi interactor.OrderInteractor) OrderController {
	return &orderController{
		OrderInteractor: oi,
	}
}

func (oc *orderController) CreateOrder(order *entities.Orders) (*entities.Orders, error) {
	return oc.OrderInteractor.CreateOrder(order)
}

func (oc *orderController) ReadOrderById(order_id string) (*entities.Orders, error) {
	return oc.OrderInteractor.ReadOrderById(order_id)
}

func (oc *orderController) UpdateOrder(order *entities.Orders, ID string) (*entities.Orders, error) {
	return oc.OrderInteractor.UpdateOrder(order, ID)
}

func (oc *orderController) DeleteOrder(order *entities.Orders, ID string) (*entities.Orders, error) {
	return oc.OrderInteractor.DeleteOrder(order, ID)
}
