package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type authPresenter struct {
}

func NewAuthPresenter() presenter.AuthPresenter {
	return &authPresenter{}
}

func (ap *authPresenter) AuthResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"error":  nil,
	}
}

func (ap *authPresenter) AuthErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err.Error(),
	}
}
