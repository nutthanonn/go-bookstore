package registry

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Book:      r.NewBookController(),
		Employee:  r.NewEmployeeController(),
		Customer:  r.NewCustomerController(),
		Inventory: r.NewInventoryController(),
		Order:     r.NewOrderController(),
	}
}
