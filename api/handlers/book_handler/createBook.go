package bookhandler

// Title        string    `gorm:"not null" json:"title"`
// Author       string    `gorm:"not null" json:"author"`
// Publish_year string    `gorm:"not null" json:"publish_year"`
// Price        float32   `gorm:"not null" json:"price"`

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/api/handlers"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	"github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
)

func (h *bookHandler) CreateBook(ca controller.AppController) fiber.Handler {
	return func(c *fiber.Ctx) error {

		fields := []string{"Title", "Author", "Publish_year", "Price"}

		var book entities.Books
		p := presenter.NewBookPresenter()
		handle := handlers.NewAppHandler()

		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(err))
		}

		if err := handle.ValidateFields(fields, book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(p.BookErrorResponse(err))
		}

		createdBook, err := ca.Book.CreateBook(&book)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(p.BookErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(p.BookCreateResponse(createdBook))
	}
}
