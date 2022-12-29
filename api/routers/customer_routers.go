package routers

import (
	"github.com/gofiber/fiber/v2"
	customerHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/customer_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func CustomerHandler(app fiber.Router, ca controller.AppController) {
	ch := customerHandler.NewCustomerHandlers()

	app.Get("/customer", ch.GetCustomer(ca))
	app.Post("/customer", ch.CreateCustomer(ca))
	app.Get("/customer/:id", ch.GetCustomerById(ca))
	app.Put("/customer/:id", ch.UpdateCustomer(ca))
	app.Delete("/customer/:id", ch.DeleteCustomer(ca))
}
