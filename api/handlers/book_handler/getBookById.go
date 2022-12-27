package bookhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (b *bookHandler) GetBookById(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID := c.Params("id")
		p := presenter.NewBookPresenter()
		if ID == "" || len(ID) > 24 {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(fiber.ErrBadRequest))
		}

		book, err := ca.Book.ReadBookByID(ID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.BookErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.BookSuccessResponse(book))
	}
}
