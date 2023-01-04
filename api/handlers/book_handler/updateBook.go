package bookhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	"gorm.io/gorm"
)

func (h *bookHandler) UpdateBook(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var book entities.Books
		id := c.Params("id")
		p := presenter.NewBookPresenter()
		handle := handlers.NewAppHandler()

		if !handle.IsValidUUID(id) {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(errors.New("invalid params id")))
		}

		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(err))
		}

		_, err := ca.Book.UpdateBook(&book, id)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(p.BookErrorResponse(errors.New("record not found")))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(p.BookErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.BookUpdateResponse())
	}
}
