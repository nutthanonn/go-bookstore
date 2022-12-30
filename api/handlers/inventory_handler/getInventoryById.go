package inventoryhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *inventoryHandler) GetInventoryById(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		p := presenter.NewInventoryPresenter()

		if !h.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.InventoryErrorResponse(errors.New("invalid params id")))
		}

		var inventory *entities.Inventories

		inventory, err := ca.Inventory.ReadInventoryById(id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.InventoryErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.InventorySuccessResponse(inventory))
	}
}
