package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type OrderPresenter interface {
	OrderSuccessResponse(order *entities.Orders) *fiber.Map
	OrderErrorResponse(err error) *fiber.Map
	OrderDeleteResponse() *fiber.Map
}
