package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type salePresenter struct {
}

func NewSalePresenter() presenter.SalePresenter {
	return &salePresenter{}
}

func (sp *salePresenter) SaleSuccessResponse(sales *entities.Sales) *fiber.Map {
	return &fiber.Map{
		"data":   sales,
		"status": true,
		"error":  nil,
	}
}

func (sp *salePresenter) SaleErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":   nil,
		"status": false,
		"error":  err.Error(),
	}

}
