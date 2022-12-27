package bookhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (b *bookHandler) UpdateBook(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var book entities.Books
		id := c.Params("id")
		p := presenter.NewBookPresenter()

		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(errors.New("id is required")))
		}

		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(err))
		}

		updatedBook, err := ca.Book.UpdateBook(&book, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.BookErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.BookSuccessResponse(updatedBook))
	}
}
