package routers

import (
	"github.com/gofiber/fiber/v2"
	bookHandler "github.com/nutthanonn/go-clean-architecture/api/handlers/book_handler"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
)

func BookRouter(app fiber.Router, ca controller.AppController) {
	bh := bookHandler.NewBookHandler()

	app.Get("/books", bh.GetBook(ca))
	app.Get("/books/:id", bh.GetBookById(ca))
	app.Post("/books", bh.CreateBook(ca))
	app.Put("/books/:id", bh.UpdateBook(ca))
	app.Delete("/books", bh.DeleteBook(ca))
}
