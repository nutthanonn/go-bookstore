package bookhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (b *bookHandler) GetBookById(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		p := presenter.NewBookPresenter()

		if !b.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(errors.New("invalid params id")))
		}

		book, err := ca.Book.ReadBookByID(id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.BookErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.BookSuccessResponse(book))
	}
}
