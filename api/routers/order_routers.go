package routers

import (
	"github.com/gofiber/fiber/v2"
	orderHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/order_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func OrderRouter(app fiber.Router, ca controller.AppController) {
	oh := orderHandler.NewOrderHandler()

	app.Post("/order", oh.CreateOrder(ca))
	app.Get("/order/:id", oh.GetOrderById(ca))
	app.Delete("/order/:id", oh.DeleteOrder(ca))
}
