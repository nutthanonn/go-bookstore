package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type SalePresenter interface {
	SaleSuccessResponse(sales *entities.Sales) *fiber.Map
	SaleErrorResponse(err error) *fiber.Map
}
