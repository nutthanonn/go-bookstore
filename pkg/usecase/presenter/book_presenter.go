package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type BookPresenter interface {
	BookCreateResponse(data *entities.Books) *fiber.Map
	BookSuccessResponse(data *entities.Books) *fiber.Map
	BooksSuccessResponse(data []*entities.Books) *fiber.Map
	BookErrorResponse(err error) *fiber.Map
	BookDeleteResponse() *fiber.Map
	BookUpdateResponse() *fiber.Map
}
