package inventoryhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	"gorm.io/gorm"
)

func (h *inventoryHandler) UpdateInventory(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		p := presenter.NewInventoryPresenter()
		handle := handlers.NewAppHandler()

		if !handle.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.InventoryErrorResponse(errors.New("invalid params id")))
		}

		var inventory *entities.Inventories

		if err := c.BodyParser(&inventory); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.InventoryErrorResponse(err))
		}

		_, err := ca.Inventory.UpdateInventory(inventory, id)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(p.InventoryErrorResponse(errors.New("record not found")))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(p.InventoryErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.InventoryUpdateResponse())
	}
}
