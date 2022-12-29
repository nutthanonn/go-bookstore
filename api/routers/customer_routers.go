package routers

import (
	"github.com/gofiber/fiber/v2"
	customerHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/customer_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func CustomerHandler(app fiber.Router, ca controller.AppController) {
	ch := customerHandler.NewCustomerHandlers()

	app.Get("/customers", ch.GetCustomer(ca))
}
