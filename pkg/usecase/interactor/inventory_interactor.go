package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type inventoryInteractor struct {
	InventoryPresenter  presenter.InventoryPresenter
	InventoryRepository repository.Inventory
}

func NewInventoryInteractor(ip presenter.InventoryPresenter, ir repository.Inventory) repository.Inventory {
	return &inventoryInteractor{
		InventoryPresenter:  ip,
		InventoryRepository: ir,
	}
}

func (ii *inventoryInteractor) ReadInventoryById(book_id string) (*entities.Inventories, error) {
	inventory, err := ii.InventoryRepository.ReadInventoryById(book_id)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (ii *inventoryInteractor) UpdateInventory(inventory *entities.Inventories, ID string) (*entities.Inventories, error) {
	updatedInventory, err := ii.InventoryRepository.UpdateInventory(inventory, ID)
	if err != nil {
		return nil, err
	}

	return updatedInventory, nil
}
