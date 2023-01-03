package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type Inventory struct {
	Book_id  uuid.UUID `json:"book_id"`
	Quantity int       `json:"quantity"`
}

type inventoryPresenter struct {
}

func NewInventoryPresenter() presenter.InventoryPresenter {
	return &inventoryPresenter{}
}

func (ip *inventoryPresenter) InventorySuccessResponse(data *entities.Inventories) *fiber.Map {
	var response = &Inventory{
		Book_id:  data.Book_id,
		Quantity: data.Quantity,
	}

	return &fiber.Map{
		"data":   response,
		"status": true,
		"error":  nil,
	}
}

func (ip *inventoryPresenter) InventoryUpdateResponse() *fiber.Map {
	return &fiber.Map{
		"data":   nil,
		"status": true,
		"error":  nil,
	}
}

func (ip *inventoryPresenter) InventoryErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"error":  err.Error(),
		"status": false,
	}
}
