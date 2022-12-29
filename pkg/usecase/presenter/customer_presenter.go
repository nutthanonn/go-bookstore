package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type CustomerPresenter interface {
	CustomerSuccessResponse(*entities.Customers) *fiber.Map
	CustomersSuccessResponse([]*entities.Customers) *fiber.Map
	CustomerErrorResponse(err error) *fiber.Map
	CustomerDeleteSuccessResponse() *fiber.Map
}
