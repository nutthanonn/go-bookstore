package routers

import (
	"github.com/gofiber/fiber/v2"
	inventoryHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/inventory_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func InventoryRouter(app fiber.Router, ca controller.AppController) {
	ih := inventoryHandler.NewInventoryHandler()

	app.Get("/inventory/:id", ih.GetInventoryById(ca))
	app.Put("/inventory/:id", ih.UpdateInventory(ca))
}
