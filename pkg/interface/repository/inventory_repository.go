package repository

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type inventoryRespoitory struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) repository.InventoryRepository {
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

	if err := ir.db.Model(&inventory).Where("employee_id = ?", ID).Updates(inventory).Error; err != nil {
		return nil, err
	}

	return inventory, nil
}
