package routers

import (
	"github.com/gofiber/fiber/v2"
	bookHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/book_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func BookRouter(app fiber.Router, ca controller.AppController) {
	bh := bookHandler.NewBookHandler()

	app.Get("/book", bh.GetBook(ca))
	app.Get("/book/:id", bh.GetBookById(ca))
	app.Post("/book", bh.CreateBook(ca))
	app.Put("/book/:id", bh.UpdateBook(ca))
	app.Delete("/book/:id", bh.DeleteBook(ca))
}
