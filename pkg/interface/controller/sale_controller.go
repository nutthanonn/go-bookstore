package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type saleController struct {
	SaleInteractor interactor.SalesInteractor
}

type SaleController interface {
	CreateSale(data *entities.Sales) (*entities.Sales, error)
}

func NewSaleController(si interactor.SalesInteractor) SaleController {
	return &saleController{SaleInteractor: si}
}

func (sc *saleController) CreateSale(data *entities.Sales) (*entities.Sales, error) {
	return sc.SaleInteractor.CreateSales(data)
}
