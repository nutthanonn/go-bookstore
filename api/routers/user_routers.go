package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func UserRouter(app fiber.Router, co controller.AppController) {
	app.Post("/users", handlers.CreateUser(co))
	app.Get("/users", handlers.GetUsers(co))
	app.Delete("/users/:id", handlers.DeleteUser(co))
	app.Put("/users", handlers.UpdateUser(co))
}
