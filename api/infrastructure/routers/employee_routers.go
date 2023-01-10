package routers

import (
	"github.com/gofiber/fiber/v2"
	employeeHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/employee_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func EmployeeRouter(app fiber.Router, ca controller.AppController) {
	eh := employeeHandler.NewEmployeeHandlers()

	app.Get("/employee", eh.GetEmployee(ca))
	app.Get("/employee/:id", eh.GetEmployeeById(ca))
	app.Post("/employee", eh.CreateEmployee(ca))
	app.Put("/employee/:id", eh.UpdateEmployee(ca))
	app.Delete("/employee/:id", eh.DeleteEmployee(ca))
}
