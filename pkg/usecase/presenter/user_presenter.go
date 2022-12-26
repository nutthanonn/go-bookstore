package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type UserPresenter interface {
	UserSuccessResponse(data *entities.User) *fiber.Map
	UsersSuccessResponse(data *[]entities.User) *fiber.Map
	UserErrorResponse(err error) *fiber.Map
	UserDeleteSuccessResponse() *fiber.Map
}
