package bookhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (b *bookHandler) CreateBook(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var book entities.Books
		p := presenter.NewBookPresenter()

		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(err))
		}

		createdBook, err := ca.Book.CreateBook(&book)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.BookErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.BookSuccessResponse(createdBook))
	}
}
