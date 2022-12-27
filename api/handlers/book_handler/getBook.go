package bookhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (b *bookHandler) GetBook(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		books, err := ca.Book.ReadBook()
		p := presenter.NewBookPresenter()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.BookErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.BooksSuccessResponse(books))
	}
}
