package routers

import (
	"github.com/gofiber/fiber/v2"
	bookhandler "github.com/nutthanonn/go-clean-architecture/api/handlers/book_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func BookRouter(app fiber.Router, ca controller.AppController) {
	bh := bookhandler.NewBookHandler()
	app.Get("/books", bh.GetBook(ca))
}
