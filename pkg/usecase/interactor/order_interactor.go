package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type orderInteractor struct {
	OrderPresenter  presenter.OrderPresenter
	OrderRepository repository.OrderRepository
}

type OrderInteractor interface {
	CreateOrder(order *entities.Orders) (*entities.Orders, error)
	ReadOrderById(order_id string) (*entities.Orders, error)
	UpdateOrder(order *entities.Orders, ID string) (*entities.Orders, error)
	DeleteOrder(ID string) error
}

func NewOrderInteractor(op presenter.OrderPresenter, or repository.OrderRepository) OrderInteractor {
	return &orderInteractor{
		OrderPresenter:  op,
		OrderRepository: or,
	}
}

func (oi *orderInteractor) CreateOrder(order *entities.Orders) (*entities.Orders, error) {
	order, err := oi.OrderRepository.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (oi *orderInteractor) ReadOrderById(order_id string) (*entities.Orders, error) {
	order, err := oi.OrderRepository.ReadOrderById(order_id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (oi *orderInteractor) UpdateOrder(order *entities.Orders, ID string) (*entities.Orders, error) {
	order, err := oi.OrderRepository.UpdateOrder(order, ID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (oi *orderInteractor) DeleteOrder(ID string) error {

	if err := oi.OrderRepository.DeleteOrder(ID); err != nil {
		return err
	}

	return nil
}
