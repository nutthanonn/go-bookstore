package repository

import (
	"errors"
	"time"

	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type inventoryRespoitory struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) repository.Inventory {
	return &inventoryRespoitory{db}
}

func (ir *inventoryRespoitory) ReadInventoryById(book_id string) (*entities.Inventories, error) {
	var inventory *entities.Inventories
	if err := ir.db.Preload("books").First(&inventory, "book_id = ?", book_id).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}

func (ir *inventoryRespoitory) UpdateInventory(inventory *entities.Inventories, ID string) (*entities.Inventories, error) {
	var oldInventory *entities.Inventories

	if err := ir.db.First(&oldInventory, "book_id = ?", ID).Error; err != nil {
		return nil, err
	}

	if oldInventory == nil {
		return nil, errors.New("inventory not found")
	}

	inventory.Book_id = oldInventory.Book_id
	inventory.Created_at = oldInventory.Created_at
	inventory.Updated_at = time.Now()

	if err := ir.db.Save(&inventory).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}
