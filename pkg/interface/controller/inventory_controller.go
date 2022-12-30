package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type inventoryController struct {
	InventoryInteractor interactor.InventoryInteractor
}

type InventoryController interface {
	ReadInventoryById(book_id string) (*entities.Inventories, error)
	UpdateInventory(inventory *entities.Inventories, ID string) (*entities.Inventories, error)
}

func NewInventoryController(ii interactor.InventoryInteractor) InventoryController {
	return &inventoryController{
		InventoryInteractor: ii,
	}
}

func (ic *inventoryController) ReadInventoryById(book_id string) (*entities.Inventories, error) {
	return ic.InventoryInteractor.ReadInventoryById(book_id)
}

func (ic *inventoryController) UpdateInventory(inventory *entities.Inventories, ID string) (*entities.Inventories, error) {
	return ic.InventoryInteractor.UpdateInventory(inventory, ID)
}
