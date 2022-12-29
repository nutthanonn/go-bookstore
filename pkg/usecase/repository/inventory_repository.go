package repository

import "github.com/nutthanonn/go-clean-architecture/pkg/entities"

type Inventory interface {
	ReadInventoryById(book_id string) (*entities.Inventories, error)
	UpdateInventory(inventory *entities.Inventories, ID string) (*entities.Inventories, error)
}
