package repository

import "github.com/nutthanonn/go-clean-architecture/pkg/entities"

type OrderRepository interface {
	CreateOrder(order *entities.Orders) (*entities.Orders, error)
	ReadOrderById(order_id string) (*entities.Orders, error)
	UpdateOrder(order *entities.Orders, ID string) (*entities.Orders, error)
	DeleteOrder(ID string) error
}
