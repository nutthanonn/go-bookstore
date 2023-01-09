package repository

import (
	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db}
}

func (or *orderRepository) CreateOrder(order *entities.Orders) (*entities.Orders, error) {
	id := uuid.New()
	order.Order_id = id

	if err := or.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (or *orderRepository) ReadOrderById(order_id string) (*entities.Orders, error) {
	var order entities.Orders
	if err := or.db.Preload("OrderDetails").Where("order_id = ?", order_id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (or *orderRepository) UpdateOrder(order *entities.Orders, ID string) (*entities.Orders, error) {
	if err := or.db.Model(&order).Where("order_id = ?", ID).Updates(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (or *orderRepository) DeleteOrder(ID string) error {

	or.db.Begin()

	if err := or.db.Where("order_id = ?", ID).Delete(&entities.OrderDetails{}).Error; err != nil {
		return err
	}

	if err := or.db.Where("order_id = ?", ID).Delete(&entities.Orders{}).Error; err != nil {
		return err
	}

	or.db.Commit()
	return nil
}
