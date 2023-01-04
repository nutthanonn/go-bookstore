package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type InventoryPresenter interface {
	InventorySuccessResponse(data *entities.Inventories) *fiber.Map
	InventoryUpdateResponse() *fiber.Map
	InventoryErrorResponse(err error) *fiber.Map
}
